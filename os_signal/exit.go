package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	_ = runtime.GOMAXPROCS(1)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGABRT, syscall.SIGTERM)

	//
	//time.AfterFunc(time.Second*3, func() {
	//	log.Printf("server is running 3mins, you can shutdown by SIGREEM!\n")
	//})
	//<-signalChan

	//
	num := int64(0)
	for i := 0; i < 1E3; i++ {
		time.AfterFunc(time.Second*2, func() {
			atomic.AddInt64(&num, 1)
		})
	}
	t := 0
	for i := 0; i < 1E10; i++ {
		t++
	}
	_ = t
	<-signalChan

	fmt.Printf("num: %v\n", num)

}
