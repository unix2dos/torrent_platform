package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"torrent_platform/cmd/ks-download/client"
	"torrent_platform/cmd/ks-download/engine"
)

type Server struct {
	engine *engine.Engine
	client *client.Client
}

func (s *Server) Run(addr string) {

	s.engine = engine.New()
	s.engine.Configure()

	s.client = client.New()

	go func() {
		for {
			s.Download()
			time.Sleep(time.Minute)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.engine.ShowDebug(w)
	})
	http.ListenAndServe(addr, mux)
}

func (s *Server) Download() {

	res := s.client.GetHash()

	for _, hash := range res {

		fmt.Println("-------------------------hash download", hash)
		torrent, err := s.engine.AddMagnet(hash)
		if err != nil {
			log.Printf("add magnet error %s\n", err)
		}

		go func() {
			<-torrent.GotInfo()
			// tTorrent.Info().Name = ""//TODO: 暂时当前目录
			torrent.DownloadAll()
		}()
	}
}
