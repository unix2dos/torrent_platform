package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Path struct {
	Path string `json:"path"`
}

func HandlerReceive() {

	router := gin.New()

	router.GET("/", handlerDebugInfo)
	path := router.Group("/path/v1")
	{
		path.PUT("/add", handlerAddPath)
		path.DELETE("/del", handlerDelPath)
	}

	router.Run(httpAddr)
}

func handlerDebugInfo(c *gin.Context) {
	torrentClient.WriteStatus(c.Writer)
}

func handlerAddPath(c *gin.Context) {

	var res Path
	err := c.BindJSON(&res)
	if err != nil {
		return
	}

	err = AddFileSeed(res.Path)
	c.JSON(200, gin.H{
		"err":  fmt.Sprintf("%v", err),
		"path": res.Path,
	})

}

func handlerDelPath(c *gin.Context) {
	var res Path
	err := c.BindJSON(&res)
	if err != nil {
		return
	}

	err = DelFileSeed(res.Path)

	c.JSON(200, gin.H{
		"err":  fmt.Sprintf("%v", err),
		"path": res.Path,
	})
}
