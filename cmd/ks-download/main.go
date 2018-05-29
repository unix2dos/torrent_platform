package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/anacrolix/torrent"

	"torrent_platform/base"
)

var (
	path     string
	httpAddr string

	torrentClient *torrent.Client
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.StringVar(&path, "path", "", "upload path")
	flag.StringVar(&httpAddr, "httpAddr", ":16184", "http addr")
}

func main() {

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

	//http listen
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

	for {
		DownloadFile()
		time.Sleep(time.Minute)
	}

	select {}
}

func DownloadFile() {

	res := GetHash()
	for _, hash := range res {

		fmt.Println("-------------------------hash download", hash)

		tTorrent, err := torrentClient.AddMagnet("magnet:?xt=urn:btih:" + hash)
		if err != nil {
			log.Printf("add magnet error %s\n", err)
		}
		go func() {
			<-tTorrent.GotInfo()
			// tTorrent.Info().Name = ""//TODO: 暂时当前目录
			tTorrent.DownloadAll()
		}()
	}
}
