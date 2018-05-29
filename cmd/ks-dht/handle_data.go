package main

import (
	"github.com/gin-gonic/gin"

	"torrent_platform/base"
)

func ListenData() {
	router := gin.New()
	router.PUT("/hash", putHash)
	router.DELETE("/hash", delHash)
	router.GET("/hash", getHash)
	router.Run(":26181")
}

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
