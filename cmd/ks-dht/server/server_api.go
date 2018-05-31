package server

import (
	"github.com/gin-gonic/gin"

	"torrent_platform/cmd/ks-dht/util"
)

var (
	hashSlice = make([]string, 0)
)

func (s *Server) handleGetHash(c *gin.Context) {
	c.JSON(200, hashSlice)
}

func (s *Server) handleAddHash(c *gin.Context) {

	var args HashArgs
	c.Bind(&args)

	var exist bool
	for _, v := range hashSlice {
		if v == args.Hash {
			exist = true
			break
		}
	}
	if !exist {
		hashSlice = append(hashSlice, args.Hash)
	}

	c.JSON(200, gin.H{"msg": "ok"})
}

func (s *Server) handleDelHash(c *gin.Context) {

	var args HashArgs
	c.Bind(&args)

	for k, v := range hashSlice {
		if v == args.Hash {
			hashSlice = util.SliceRemove(hashSlice, k)
			break
		}
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
