package rtm

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/gorilla/websocket"
	pdebug "github.com/lestrrat/go-pdebug"
	"github.com/lestrrat/go-slack"
	"github.com/pkg/errors"
)

func New(cl *slack.Client) *Client {
	return &Client{
		client:   cl,
		eventsCh: make(chan *Event),
	}
}

func (c *Client) Events() <-chan *Event {
	return c.eventsCh
}

// Run starts the RTM run loop.
func (c *Client) Run(octx context.Context) error {
	octxwc, cancel := context.WithCancel(octx)
	defer cancel()

	ctx := newRtmCtx(octxwc, c.eventsCh)
	go ctx.run()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		ctx.emit(&Event{typ: ClientConnectingEventType})

		var conn *websocket.Conn

		strategy := backoff.NewExponentialBackOff()
		strategy.InitialInterval = 100 * time.Millisecond
		strategy.MaxInterval = 5 * time.Second
		strategy.MaxElapsedTime = 0
		err := backoff.Retry(func() error {
			res, err := c.client.RTM().Start().Do(ctx)
			if err != nil {
				log.Printf("failed to start RTM sesson: %s", err)
				return err
			}
			conn, _, err = websocket.DefaultDialer.Dial(res.URL, nil)
			if err != nil {
				log.Printf("failed to dial to websocket: %s", err)
				return err
			}
			return nil
		}, backoff.WithContext(strategy, ctx))

		if err != nil {
			return errors.Wrap(err, `failed to connect to RTM endpoint`)
		}

		ctx.handleConn(conn)
		// we get here if we manually canceled the context
		// of if the websocket ReadMessage returned an error
		ctx.emit(&Event{typ: ClientDisconnectedEventType})

	}

	return nil
}

func (ctx *rtmCtx) handleConn(conn *websocket.Conn) {
	defer conn.Close()

	in := make(chan []byte)

	// This goroutine is responsible for reading from the
	// websocket connection. It's separated because the
	// ReadMessage() operation is blocking.
	go func(ch chan []byte, conn *websocket.Conn) {
		defer close(ch)

		for {
			typ, data, err := conn.ReadMessage()
			if err != nil {
				// There was an error. we need to bail out
				if pdebug.Enabled {
					pdebug.Printf("error while reading message from websocket: %s", err)
				}
				return
			}

			// we only understand text messages
			if typ != websocket.TextMessage {
				if pdebug.Enabled {
					pdebug.Printf("received websocket message, but it is not a text payload. refusing to process")
				}
				continue
			}
			if pdebug.Enabled {
				pdebug.Printf("forwarding new websocket message")
			}
			ch <- data
		}
	}(in, conn)

	for {
		select {
		case <-ctx.Done():
			return
		case payload, ok := <-in:
			if !ok {
				if pdebug.Enabled {
					pdebug.Printf("websocket proxy: detected incoming channel close.")
				}
				// if the channel is closed, we probably had some
				// problems in the ReadMessage proxy. bail out
				return
			}

			if pdebug.Enabled {
				pdebug.Printf("websocket proxy: received raw payload: %s", payload)
			}

			var event Event
			if err := json.Unmarshal(payload, &event); err != nil {
				if pdebug.Enabled {
					pdebug.Printf("websocket proxy: failed to unmarshal payload: %s", err)
				}
			}

			ctx.inbuf <- &event
		}
	}
}

type rtmCtx struct {
	context.Context
	inbuf        chan *Event
	outbuf       chan<- *Event
	writeTimeout time.Duration
}

func newRtmCtx(octx context.Context, outch chan<- *Event) *rtmCtx {
	return &rtmCtx{
		Context:      octx,
		inbuf:        make(chan *Event),
		outbuf:       outch,
		writeTimeout: 500 * time.Millisecond,
	}
}

// Attempt to write to the outgoing channel, within the
// alloted time frame.
func (ctx *rtmCtx) trywrite(e *Event) error {
	tctx, cancel := context.WithTimeout(ctx, ctx.writeTimeout)
	defer cancel()

	select {
	case <-tctx.Done():
		switch err := tctx.Err(); err {
		case context.DeadlineExceeded:
			return errors.New("write timeout")
		default:
			return err
		}
	case ctx.outbuf <- e:
		return nil
	}

	return errors.New("unreachable")
}

// The point of this loop is to ensure the writer (the loop receiving
// events from the websocket connection) can safely write the events
// to a channel without worrying about blocking.
//
// Inside this loop, we read from the channel receiving the events,
// and we either write to the consumer channel, or buffer in our
// in memory queue (list) for later consumption
func (ctx *rtmCtx) run() {
	defer close(ctx.outbuf) // make sure the reader of Events() gets notified

	periodic := time.NewTicker(time.Second)
	defer periodic.Stop()

	var events []*Event
	for {
		select {
		case <-ctx.Done():
			return
		case e := <-ctx.inbuf:
			events = append(events, e)
		case <-periodic.C:
			// attempt to flush the buffer periodically.
		}

		// events should only contain more than one item if we
		// failed to write to the outgoing channel within the
		// allotted time
		for len(events) > 0 {
			e := events[0]
			// Try writing. if we fail, bail out of this write loop
			if err := ctx.trywrite(e); err != nil {
				break
			}
			// if we were successful, pop the current one and try the next one
			events = events[1:]
		}

		// shink the slice if we're too big
		if l := len(events); l > 16 && cap(events) > 2*l {
			events = append([]*Event(nil), events...)
		}
	}
}

// emit sends the event e to a channel. This method doesn't "fail" to
// write because we expect the the proxy loop in run() to read these
// requests as quickly as possible under normal circumstances
func (ctx *rtmCtx) emit(e *Event) {
	select {
	case <-ctx.Done():
		return
	case ctx.inbuf <- e:
		return
	}
}
