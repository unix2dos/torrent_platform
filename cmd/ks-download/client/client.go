package client

import "torrent_platform/base"

type Client struct {
	*base.Client
}

func New() *Client {
	return &Client{base.NewClient("207.246.80.69:26180")}
	// return &Client{base.NewClient("172.24.120.28:26180")}

}

func (c *Client) GetHash() (res []string) {
	c.DoGet(c.UrlFor(base.UtHash), &res)
	return
}
