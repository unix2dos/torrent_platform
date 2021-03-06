package client

import (
	"bytes"
	"encoding/json"

	"torrent_platform/base"
)

type Client struct {
	*base.Client
}

func New() *Client {
	return &Client{base.NewClient("207.246.80.69:26180")}
	// return &Client{base.NewClient("172.24.120.28:26180")}
}

func (c *Client) AddHash(hash string) error {

	h := HashArgs{Hash: hash}
	buf, err := json.Marshal(h)
	if err != nil {
		return err
	}

	return c.DoPut(c.UrlFor(base.UtHash), "application/json", bytes.NewBuffer(buf))
}

func (c *Client) DelHash(hash string) error {

	h := HashArgs{Hash: hash}
	buf, err := json.Marshal(h)
	if err != nil {
		return err
	}

	return c.DoDelete(c.UrlFor(base.UtHash), "application/json", bytes.NewBuffer(buf))
}
