package main

import (
	"torrent_platform/base"
)

var (
	httpClient = base.NewClient()
)

func GetHash() (res []string) {
	httpClient.DoGet("http://127.0.0.1:26181/hash", &res)
	return
}
