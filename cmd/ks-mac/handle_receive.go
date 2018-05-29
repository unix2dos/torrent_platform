package main

import (
	"github.com/gin-gonic/gin"
)

type Path struct {
	Path string `json:"path"`
}

func HandlerReceive() {

	router := gin.New()
	router.GET("/", handlerDebugInfo)
	router.PUT("/addPath", handlerAddPath)
	router.DELETE("/delPath", handlerDelPath)
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

	c.String(200, "path:%s \nerr=%v\n", res.Path, err)
	AddFileSeed(res.Path)

}

func handlerDelPath(c *gin.Context) {
	var res Path
	err := c.BindJSON(&res)
	if err != nil {
		return
	}

	c.String(200, "path:%s \nerr=%v\n", res.Path, err)
	DelFileSeed(res.Path)
}
