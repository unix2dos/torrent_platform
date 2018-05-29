package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/anacrolix/torrent"

	"torrent_platform/base"
)

var (
	httpAddr string

	torrentClient *torrent.Client
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.StringVar(&httpAddr, "httpAddr", ":16183", "http addr")
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

	//listen path
	go func() {
		ListenPath()
	}()

	//http debug
	listener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer listener.Close()
		log.Printf("error serving http on envpprof listener: %s", http.Serve(listener, nil))
	}()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		torrentClient.WriteStatus(w)
	})

	select {}
}
