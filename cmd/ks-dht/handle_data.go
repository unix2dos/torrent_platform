package main

import (
	"github.com/gin-gonic/gin"

	"torrent_platform/base"
)

var (
	hashSlice = make([]string, 0)
)

func handleGetHash(c *gin.Context) {

	c.JSON(200, hashSlice)
}

func handleAddHash(c *gin.Context) {
	var hash base.Hash
	c.BindJSON(&hash)

	hashSlice = append(hashSlice, hash.Hash)

	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func handleDelHash(c *gin.Context) {

	c.JSON(200, gin.H{
		"msg": "ok",
	})
}
