package main

import (
	"bytes"
	"encoding/json"

	"torrent_platform/base"
)

var (
	httpClient = base.NewClient()
)

func AddHash(hash string) error {

	h := base.Hash{Hash: hash}
	buf, err := json.Marshal(h)
	if err != nil {
		return err
	}

	return httpClient.DoPut(httpClient.UrlFor(base.UtHash), "application/json", bytes.NewBuffer(buf))
}

func DelHash(hash string) error {

	h := base.Hash{Hash: hash}
	buf, err := json.Marshal(h)
	if err != nil {
		return err
	}

	return httpClient.DoDelete(httpClient.UrlFor(base.UtHash), "application/json", bytes.NewBuffer(buf))
}
