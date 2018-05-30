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

	c.JSON(200, gin.H{"msg": "ok"})
}

func handleDelHash(c *gin.Context) {

	var hash base.Hash
	c.BindJSON(&hash)

	for k, v := range hashSlice {
		if v == hash.Hash {
			hashSlice = remove(hashSlice, k)
			break
		}
	}

	c.JSON(200, gin.H{"msg": "ok"})
}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
