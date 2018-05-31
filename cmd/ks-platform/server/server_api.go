package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleDebugInfo(c *gin.Context) {
	s.engine.ShowDebug(c.Writer)
}

func (s *Server) handleAddPath(c *gin.Context) {

	var args PathArgs
	err := c.Bind(&args)
	if err != nil {
		return
	}

	_, err = os.Stat(args.Path)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"msg": "path not exist", "ret": "1"})
		return
	}

	msg := "ok"
	infohash, err := s.engine.AddFileSeed(args.Path)
	if err != nil {
		msg = fmt.Sprintf("%v", err)
	}

	err = s.client.AddHash(infohash)
	if err != nil {
		msg = fmt.Sprintf("%v", err)
	}

	c.JSON(200, gin.H{"msg": msg, "ret": "0"})

}

func (s *Server) handleDelPath(c *gin.Context) {

	var args PathArgs
	err := c.Bind(&args)
	if err != nil {
		return
	}

	_, err = os.Stat(args.Path)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"msg": "path not exist", "ret": "1"})
		return
	}

	msg := "ok"
	infohash, err := s.engine.DelFileSeed(args.Path)
	if err != nil {
		msg = fmt.Sprintf("%v", err)
	}

	err = s.client.DelHash(infohash) //TODO:要不要删除, 因为别人有可能也在做种
	if err != nil {
		msg = fmt.Sprintf("%v", err)
	}

	c.JSON(200, gin.H{"msg": msg, "ret": "0"})
}

func (s *Server) handleListPath(c *gin.Context) {

	//TODO: list path
	c.JSON(200, gin.H{"msg": "ok", "ret": "0"})
}
