package main

import (
	"torrent_platform/base"
)

var (
	httpClient = base.NewClient()
)

func GetHash() (res []string) {
	httpClient.DoGet(httpClient.UrlFor(base.UtHash), &res)
	return
}
