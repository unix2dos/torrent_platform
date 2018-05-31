package client

import "torrent_platform/base"

type Client struct {
	*base.Client
}

func New() *Client {
	return &Client{base.NewClient()}
}

func (c *Client) GetHash() (res []string) {
	c.DoGet(c.UrlFor(base.UtHash), &res)
	return
}
