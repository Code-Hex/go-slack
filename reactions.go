package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = objects.EpochTime(0)

// ReactionsAddCall is created by ReactionsService.Add method call
type ReactionsAddCall struct {
	service     *ReactionsService
	channel     string
	file        string
	fileComment string
	name        string
	timestamp   string
}

// ReactionsGetCall is created by ReactionsService.Get method call
type ReactionsGetCall struct {
	service     *ReactionsService
	channel     string
	file        string
	fileComment string
	full        bool
	timestamp   string
}

// ReactionsListCall is created by ReactionsService.List method call
type ReactionsListCall struct {
	service *ReactionsService
	count   int
	full    bool
	page    int
	user    string
}

// ReactionsRemoveCall is created by ReactionsService.Remove method call
type ReactionsRemoveCall struct {
	service     *ReactionsService
	channel     string
	file        string
	fileComment string
	name        string
	timestamp   string
}

// Add creates a ReactionsAddCall object in preparation for accessing the reactions.add endpoint
func (s *ReactionsService) Add(name string) *ReactionsAddCall {
	var call ReactionsAddCall
	call.service = s
	call.name = name
	return &call
}

// Channel sets the value for optional channel parameter
func (c *ReactionsAddCall) Channel(channel string) *ReactionsAddCall {
	c.channel = channel
	return c
}

// File sets the value for optional file parameter
func (c *ReactionsAddCall) File(file string) *ReactionsAddCall {
	c.file = file
	return c
}

// FileComment sets the value for optional fileComment parameter
func (c *ReactionsAddCall) FileComment(fileComment string) *ReactionsAddCall {
	c.fileComment = fileComment
	return c
}

// Timestamp sets the value for optional timestamp parameter
func (c *ReactionsAddCall) Timestamp(timestamp string) *ReactionsAddCall {
	c.timestamp = timestamp
	return c
}

// Values() returns the ReactionsAddCall object as url.Values
func (c *ReactionsAddCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if len(c.channel) > 0 {
		v.Set("channel", c.channel)
	}

	if len(c.file) > 0 {
		v.Set("file", c.file)
	}

	if len(c.fileComment) > 0 {
		v.Set("fileComment", c.fileComment)
	}

	if len(c.name) <= 0 {
		return nil, errors.New(`missing required parameter name`)
	}
	v.Set("name", c.name)

	if len(c.timestamp) > 0 {
		v.Set("timestamp", c.timestamp)
	}
	return v, nil
}

// Do executes the call to access reactions.add endpoint
func (c *ReactionsAddCall) Do(ctx context.Context) error {
	const endpoint = "reactions.add"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		SlackResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to reactions.add`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// Get creates a ReactionsGetCall object in preparation for accessing the reactions.get endpoint
func (s *ReactionsService) Get() *ReactionsGetCall {
	var call ReactionsGetCall
	call.service = s
	return &call
}

// Channel sets the value for optional channel parameter
func (c *ReactionsGetCall) Channel(channel string) *ReactionsGetCall {
	c.channel = channel
	return c
}

// File sets the value for optional file parameter
func (c *ReactionsGetCall) File(file string) *ReactionsGetCall {
	c.file = file
	return c
}

// FileComment sets the value for optional fileComment parameter
func (c *ReactionsGetCall) FileComment(fileComment string) *ReactionsGetCall {
	c.fileComment = fileComment
	return c
}

// Full sets the value for optional full parameter
func (c *ReactionsGetCall) Full(full bool) *ReactionsGetCall {
	c.full = full
	return c
}

// Timestamp sets the value for optional timestamp parameter
func (c *ReactionsGetCall) Timestamp(timestamp string) *ReactionsGetCall {
	c.timestamp = timestamp
	return c
}

// Values() returns the ReactionsGetCall object as url.Values
func (c *ReactionsGetCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if len(c.channel) > 0 {
		v.Set("channel", c.channel)
	}

	if len(c.file) > 0 {
		v.Set("file", c.file)
	}

	if len(c.fileComment) > 0 {
		v.Set("fileComment", c.fileComment)
	}

	if c.full {
		v.Set("full", "true")
	}

	if len(c.timestamp) > 0 {
		v.Set("timestamp", c.timestamp)
	}
	return v, nil
}

// Do executes the call to access reactions.get endpoint
func (c *ReactionsGetCall) Do(ctx context.Context) (*ReactionsGetResponse, error) {
	const endpoint = "reactions.get"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*ReactionsGetResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to reactions.get`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ReactionsGetResponse, nil
}

// List creates a ReactionsListCall object in preparation for accessing the reactions.list endpoint
func (s *ReactionsService) List() *ReactionsListCall {
	var call ReactionsListCall
	call.service = s
	return &call
}

// Count sets the value for optional count parameter
func (c *ReactionsListCall) Count(count int) *ReactionsListCall {
	c.count = count
	return c
}

// Full sets the value for optional full parameter
func (c *ReactionsListCall) Full(full bool) *ReactionsListCall {
	c.full = full
	return c
}

// Page sets the value for optional page parameter
func (c *ReactionsListCall) Page(page int) *ReactionsListCall {
	c.page = page
	return c
}

// User sets the value for optional user parameter
func (c *ReactionsListCall) User(user string) *ReactionsListCall {
	c.user = user
	return c
}

// Values() returns the ReactionsListCall object as url.Values
func (c *ReactionsListCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.count > 0 {
		v.Set("count", strconv.Itoa(c.count))
	}

	if c.full {
		v.Set("full", "true")
	}

	if c.page > 0 {
		v.Set("page", strconv.Itoa(c.page))
	}

	if len(c.user) > 0 {
		v.Set("user", c.user)
	}
	return v, nil
}

// Do executes the call to access reactions.list endpoint
func (c *ReactionsListCall) Do(ctx context.Context) (*ReactionsListResponse, error) {
	const endpoint = "reactions.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*ReactionsListResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to reactions.list`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.ReactionsListResponse, nil
}

// Remove creates a ReactionsRemoveCall object in preparation for accessing the reactions.remove endpoint
func (s *ReactionsService) Remove(name string) *ReactionsRemoveCall {
	var call ReactionsRemoveCall
	call.service = s
	call.name = name
	return &call
}

// Channel sets the value for optional channel parameter
func (c *ReactionsRemoveCall) Channel(channel string) *ReactionsRemoveCall {
	c.channel = channel
	return c
}

// File sets the value for optional file parameter
func (c *ReactionsRemoveCall) File(file string) *ReactionsRemoveCall {
	c.file = file
	return c
}

// FileComment sets the value for optional fileComment parameter
func (c *ReactionsRemoveCall) FileComment(fileComment string) *ReactionsRemoveCall {
	c.fileComment = fileComment
	return c
}

// Timestamp sets the value for optional timestamp parameter
func (c *ReactionsRemoveCall) Timestamp(timestamp string) *ReactionsRemoveCall {
	c.timestamp = timestamp
	return c
}

// Values() returns the ReactionsRemoveCall object as url.Values
func (c *ReactionsRemoveCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if len(c.channel) > 0 {
		v.Set("channel", c.channel)
	}

	if len(c.file) > 0 {
		v.Set("file", c.file)
	}

	if len(c.fileComment) > 0 {
		v.Set("fileComment", c.fileComment)
	}

	if len(c.name) <= 0 {
		return nil, errors.New(`missing required parameter name`)
	}
	v.Set("name", c.name)

	if len(c.timestamp) > 0 {
		v.Set("timestamp", c.timestamp)
	}
	return v, nil
}

// Do executes the call to access reactions.remove endpoint
func (c *ReactionsRemoveCall) Do(ctx context.Context) error {
	const endpoint = "reactions.remove"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		SlackResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to reactions.remove`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}
