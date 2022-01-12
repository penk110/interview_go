package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// runtime.GC()
	go test()

	engine := gin.Default()
	err := engine.RunUnix("./gin.sock")
	if err != nil {
		panic(err)
	}

}

func test() {
	go func(options ...interface{}) {
		fmt.Printf("options: %v\n", options)
	}()
}
