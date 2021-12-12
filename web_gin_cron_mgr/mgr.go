package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/penk110/interview_go/web_gin_cron_mgr/handler"
)

var gE *gin.Engine

func main() {
	var (
		v1        *gin.RouterGroup
		cronTaskG *gin.RouterGroup
		err       error
	)
	gE = gin.Default()

	gE.GET("/", RootPageH)
	gE.GET("/favicon.ico", FaviconIcoH)
	v1 = gE.Group("/api/v1")

	cronTaskG = v1.Group("/cron")
	{
		cronTaskG.GET("", handler.GetCronTaskH().ListEntry)
		cronTaskG.DELETE("/:id", handler.GetCronTaskH().DelEntry)
	}

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
