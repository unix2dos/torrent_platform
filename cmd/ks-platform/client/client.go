package client

import (
	"bytes"
	"encoding/json"

	"torrent_platform/base"
)

type Hash struct {
	Hash string `json:"hash" binding:"required"`
}

type Client struct {
	*base.Client
}

func New() *Client {
	return &Client{base.NewClient()}
}

func (c *Client) AddHash(hash string) error {

	h := Hash{Hash: hash}
	buf, err := json.Marshal(h)
	if err != nil {
		return err
	}

	return c.DoPut(c.UrlFor(base.UtHash), "application/json", bytes.NewBuffer(buf))
}

func (c *Client) DelHash(hash string) error {

	h := Hash{Hash: hash}
	buf, err := json.Marshal(h)
	if err != nil {
		return err
	}

	return c.DoDelete(c.UrlFor(base.UtHash), "application/json", bytes.NewBuffer(buf))
}
