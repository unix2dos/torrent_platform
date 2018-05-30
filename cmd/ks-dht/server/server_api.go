package server

import (
	"github.com/gin-gonic/gin"

	"torrent_platform/base"
	"torrent_platform/cmd/ks-dht/util"
)

var (
	hashSlice = make([]string, 0)
)

func (s *Server) handleGetHash(c *gin.Context) {

	c.JSON(200, hashSlice)
}

func (s *Server) handleAddHash(c *gin.Context) {

	var hash base.Hash
	c.BindJSON(&hash)

	hashSlice = append(hashSlice, hash.Hash)
	c.JSON(200, gin.H{"msg": "ok"})
}

func (s *Server) handleDelHash(c *gin.Context) {

	var hash base.Hash
	c.BindJSON(&hash)

	for k, v := range hashSlice {
		if v == hash.Hash {
			hashSlice = util.SliceRemove(hashSlice, k)
			break
		}
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
