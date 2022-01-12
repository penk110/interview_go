package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path"
	"time"
)

func main() {
	_BasePath, _ := os.Getwd()
	_stdout := path.Join(_BasePath, "web_gin_default_writer", "stdout.log")
	_stderr := path.Join(_BasePath, "web_gin_default_writer", "stderr.log")
	stdout, err := os.Create(_stdout)
	stderr, err := os.Create(_stderr)
	if err != nil {
		fmt.Println("Open Log File Failed", err)
	}

	gin.DefaultWriter = io.MultiWriter(stdout, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(stderr, os.Stderr)

	g := gin.Default()
	fmt.Println("Init Gin log set ")
	g.GET("/hello/:name", func(c *gin.Context) {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println("fmt: ", now)
		log.Println("log: ", now)
		c.String(200, "Hello %s", c.Param("name"))
	})

	g.GET("/err", func(c *gin.Context) {
		now := time.Now().Format("2006-01-02 15:04:05")
		log.Println("[err] err handler, fmt time: ", now)
		panic(now)
		return
	})

	err = g.Run(":9000")
	if err != nil {
		panic(err)
	}
}
