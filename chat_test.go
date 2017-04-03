package slack_test

import (
	"context"
	"testing"

	"github.com/lestrrat/go-slack"
	"github.com/stretchr/testify/assert"
)

// Test message create, update, delete
func TestChatMessage(t *testing.T) {
	if !requireSlackToken(t) || !requireDMUser(t) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)

	res, err := c.Chat().PostMessage(ctx, dmUser, "hello", nil)
	if !assert.NoError(t, err, "Chat.PostMessage failed") {
		return
	}
	t.Logf("%#v", res)

	res, err = c.Chat().Update(ctx, res.Channel, res.Timestamp, "hello, world!")
	if !assert.NoError(t, err, "Chat.Update failed") {
		return
	}
	t.Logf("%#v", res)

	/*
		res, err = c.Chat().Delete(ctx, res.Channel, res.Timestamp)
		if !assert.NoError(t, err, "Chat.Delete failed") {
			return
		}
		t.Logf("%#v", res)
	*/
}

// Test me message. Note that
func TestChatMeMessage(t *testing.T) {
	if !requireSlackToken(t) || !requireDMUser(t) || !requireRealUser(t) {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := slack.New(slackToken)

	res, err := c.Chat().MeMessage(ctx, dmUser, "hello")
	if !assert.NoError(t, err, "Chat.MeMessage failed") {
		return
	}
	t.Logf("%#v", res)
}
