package main

import (
	"bytes"
	"encoding/json"

	"torrent_platform/base"
)

var (
	httpClient = base.NewClient()
)

func SendHash(hash string) error {

	h := base.Hash{Hash: hash}
	buf, err := json.Marshal(h)
	if err != nil {
		return err
	}

	return httpClient.DoPut(httpClient.UrlFor(base.UtHash), "application/json", bytes.NewBuffer(buf))
}
