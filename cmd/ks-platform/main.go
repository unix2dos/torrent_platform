package main

import (
	"flag"
	"log"

	"torrent_platform/cmd/ks-platform/server"
)

var (
	httpAddr string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.StringVar(&httpAddr, "httpAddr", ":16180", "http addr")
}

func main() {

	flag.Parse()

	server := &server.Server{}
	server.Run(httpAddr)
}
