package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{
		Addr:                   "127.0.0.1:9091",
		ConfigDir:              "./",
		ShutdownCleanup:        false,
		ReuseSocketAddrAndPort: false,
	}); err != nil {
		log.Fatalf("agent.Listen err: %v", err)
	}

	go loop()
	r := gin.Default()
	r.GET("/index", indexH)

	if err := r.Run(":9090"); err != nil {
		log.Fatal(err.Error())
	}

}

func indexH(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"code":    2000000,
		"message": "ok",
		"data":    "success",
	})
	return
}

func loop() {
	var (
		tick3s  *time.Ticker
		tick10s *time.Ticker
	)

	tick3s = time.NewTicker(time.Second * 3)
	tick10s = time.NewTicker(time.Second * 10)

	for {
		select {
		case <-tick3s.C:
			log.Println("tick3s")
		case <-tick10s.C:
			log.Println("tick10s")
		}
	}
}
