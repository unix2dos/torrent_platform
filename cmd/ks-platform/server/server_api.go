package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleDebugInfo(c *gin.Context) {
	s.engine.ShowDebug(c.Writer)
}

func (s *Server) handleAddPath(c *gin.Context) {

	var res Path
	err := c.BindJSON(&res)
	if err != nil {
		return
	}

	_, err = os.Stat(res.Path)
	if err != nil {
		return
	}

	infohash, err := s.engine.AddFileSeed(res.Path)
	if err != nil {
		return
	}

	err = s.client.AddHash(infohash)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{"msg": "ok"})

}

func (s *Server) handleDelPath(c *gin.Context) {
	var res Path
	err := c.BindJSON(&res)
	if err != nil {
		return
	}

	_, err = os.Stat(res.Path)
	if err != nil {
		return
	}

	infohash, err := s.engine.DelFileSeed(res.Path)
	if err != nil {
		return
	}

	err = s.client.DelHash(infohash) //TODO:要不要删除, 因为别人有可能也在做种
	if err != nil {
		return
	}

	c.JSON(200, gin.H{"msg": "ok"})
}

func (s *Server) handleListPath(c *gin.Context) {

	//TODO: list path
	c.JSON(200, gin.H{"msg": "ok"})
}
