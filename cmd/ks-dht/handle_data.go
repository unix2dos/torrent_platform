package main

import (
	"github.com/gin-gonic/gin"

	"torrent_platform/base"
)

var (
	hashSlice = make([]string, 0)
)

//增加
func putHash(c *gin.Context) {
	var hash base.Hash
	c.BindJSON(&hash)

	hashSlice = append(hashSlice, hash.Hash)

	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

//删除
func delHash(c *gin.Context) {

	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

//查询
func getHash(c *gin.Context) {

	c.JSON(200, hashSlice)
}
