package main

import (
	"flag"
	"log"

	"github.com/anacrolix/torrent"

	"torrent_platform/base"
)

var (
	httpAddr string

	torrentClient *torrent.Client
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.StringVar(&httpAddr, "httpAddr", ":16180", "http addr")
}

func main() {

	flag.Parse()

	//torrent client
	clientConfig := torrent.Config{
		Debug:            false,
		Seed:             true,
		DisableIPv6:      true,
		DhtStartingNodes: base.BootstrapAddrs,
		DefaultStorage:   base.NewFile("", ""),
	}
	var err error
	torrentClient, err = torrent.NewClient(&clientConfig)
	if err != nil {
		log.Printf("error creating torrent: %s", err)
	}
	defer torrentClient.Close()

	//listen http
	go func() {
		HandlerReceive()
	}()

	select {}
}
