package main

import (
	"flag"
	"fmt"
	"log"
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
		mux := http.NewServeMux()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			torrentClient.WriteStatus(w)
		})

		mux.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Query().Get("key")
			fmt.Fprintf(w, "your path is, %s!\n\n%s\n", path, AddPath(path))
		})
		http.ListenAndServe(httpAddr, mux)
	}()

	select {}
}
