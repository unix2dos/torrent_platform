package server

import (
	"github.com/gin-gonic/gin"

	"torrent_platform/cmd/ks-platform/client"
	"torrent_platform/cmd/ks-platform/engine"
)

type Server struct {
	engine *engine.Engine
	client *client.Client
}

func (s *Server) Run(addr string) {
	s.engine = engine.New()
	s.engine.Configure()

	s.client = client.New()

	r := gin.New()
	s.Route(r)
	r.Run(addr)
}

func (s *Server) Route(r *gin.Engine) {

	r.GET("/", s.handleDebugInfo)

	path := r.Group("/path/v1")
	{
		path.PUT("/add", s.handleAddPath)
		path.DELETE("/del", s.handleDelPath)
		path.GET("/list", s.handleListPath)
	}
}
