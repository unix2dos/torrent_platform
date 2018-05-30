package server

import (
	"github.com/gin-gonic/gin"

	"torrent_platform/cmd/ks-dht/util"
)

type Hash struct {
	Hash string `json:"hash" binding:"required"`
}

var (
	hashSlice = make([]string, 0)
)

func (s *Server) handleGetHash(c *gin.Context) {
	c.JSON(200, hashSlice)
}

func (s *Server) handleAddHash(c *gin.Context) {

	var hash Hash
	c.BindJSON(&hash)

	var exist bool
	for _, v := range hashSlice {
		if v == hash.Hash {
			exist = true
			break
		}
	}
	if !exist {
		hashSlice = append(hashSlice, hash.Hash)
	}

	c.JSON(200, gin.H{"msg": "ok"})
}

func (s *Server) handleDelHash(c *gin.Context) {

	var hash Hash
	c.BindJSON(&hash)

	for k, v := range hashSlice {
		if v == hash.Hash {
			hashSlice = util.SliceRemove(hashSlice, k)
			break
		}
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
