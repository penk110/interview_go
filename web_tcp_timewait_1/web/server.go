package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var gE *gin.Engine

func main() {
	var (
		err error
	)
	gE = gin.Default()

	gE.GET("/", RootPageH)
	gE.GET("/favicon.ico", FaviconIcoH)
	if err = gE.Run(":8030"); err != nil {
		log.Printf("start failed, err: %v", err.Error())
		os.Exit(2)
	}
}

func RootPageH(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 100010,
		"msg":  "ok",
		"data": "root page",
	})
	return
}

func FaviconIcoH(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 100010,
		"msg":  "ok",
		"data": "favicon.ico",
	})
	return
}
