package server

import (
	"os"

	"github.com/gin-gonic/gin"

	"fhyx/kinema/pkg/web/response"
	"torrent_platform/base"
)

func (s *Server) handleDebugInfo(c *gin.Context) {
	s.engine.ShowDebug(c.Writer)
}

func (s *Server) handleAddPath(c *gin.Context) {

	var args PathArgs
	err := c.Bind(&args)
	if err != nil {
		response.Resp400JSON(c, base.ErrParamType, err)
		return
	}

	_, err = os.Stat(args.Path)
	if err != nil {
		response.BadReqJSON(c, base.ErrorPathNotExist)
		return
	}

	infohash, err := s.engine.AddFileSeed(args.Path)
	if err != nil {
		response.BadReqJSON(c, err)
		return
	}

	err = s.client.AddHash(infohash)
	if err != nil {
		c.JSON(503, gin.H{"status": 1, "error": "系统错误"})
		return
	}

	response.SuccessJSON(c)
}

func (s *Server) handleDelPath(c *gin.Context) {

	var args PathArgs
	err := c.Bind(&args)
	if err != nil {
		response.Resp400JSON(c, base.ErrParamType, err)
		return
	}

	_, err = os.Stat(args.Path)
	if err != nil {
		response.BadReqJSON(c, base.ErrorPathNotExist)
		return
	}

	infohash, err := s.engine.DelFileSeed(args.Path)
	if err != nil {
		response.BadReqJSON(c, err)
		return
	}

	err = s.client.DelHash(infohash)
	if err != nil {
		c.JSON(503, gin.H{"status": 1, "error": "系统错误"})
		return
	}

	response.SuccessJSON(c)
}

func (s *Server) handleListPath(c *gin.Context) {

	//TODO: list path
	response.DataJSON(c, nil, 0)
}
