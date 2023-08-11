package restylog

import (
	"context"
	"crypto/tls"
	"time"

	"gopkg.in/resty.v1"
)

const defaultTimeout = 10

type Client struct {
	client *resty.Client
	tag    string
	ctx    context.Context
}

func New(ctx context.Context) *Client {
	c := &Client{
		resty.New(),
		"client",
		ctx,
	}
	c.client.SetHeader("User-Agent", "Iomob/backend")
	c.client.SetLogger(&logWriter{ctx})
	c.client.SetDebug(true)
	c.client.SetTimeout(defaultTimeout * time.Second)
	return c
}

// WithTag sets the tag for the restylog client
// The tag is used to identify the resty client in the logs in case we have multiple clients
func (c *Client) WithTag(tag string) *Client {
	c.tag = tag
	return c
}

func (c *Client) R() *Request {
	req := CreateRequest(c.ctx, c.client.R(), c.tag)
	return req
}

func (c *Client) SetHostURL(url string) *Client {
	c.client.SetHostURL(url)
	return c
}

func (c *Client) SetHeader(k, v string) *Client {
	c.client.SetHeader(k, v)
	return c
}

func (c *Client) SetTLSClientConfig(config *tls.Config) *Client {
	c.client.SetTLSClientConfig(config)
	return c
}

func (c *Client) SetBasicAuth(username, password string) *Client {
	c.client.SetBasicAuth(username, password)
	return c
}
func (c *Client) SetCloseConnection(closeConn bool) *Client {
	c.client.SetCloseConnection(closeConn)
	return c
}
