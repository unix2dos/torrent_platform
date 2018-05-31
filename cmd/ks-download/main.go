package main

import (
	"flag"
	"log"

	"torrent_platform/cmd/ks-download/server"
)

var (
	path         string
	dhtDebugAddr string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.StringVar(&path, "path", "", "upload path")
	flag.StringVar(&dhtDebugAddr, "dhtDebugAddr", ":16184", "dht debug addr")
}

func main() {

	server := &server.Server{}
	server.Run(dhtDebugAddr)
}
