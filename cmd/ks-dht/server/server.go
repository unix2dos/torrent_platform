package server

import (
	"github.com/gin-gonic/gin"

	"torrent_platform/cmd/ks-dht/dht"
)

type Server struct {
	DHT *dht.DHT
}

func (s *Server) Run(addr string) {

	go func() {
		r := gin.Default()
		s.Route(r)
		r.Run(addr)
	}()

	s.DHT.Start()
}

func (s *Server) Route(r *gin.Engine) {

	r.GET("/hash", s.handleGetHash)
	r.PUT("/hash", s.handleAddHash)
	r.DELETE("/hash", s.handleDelHash)
}
