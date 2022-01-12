package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

/*
使用time.Ticker实现定时任务
*/

// schedule task
func tack(ticker <-chan time.Time) {
	for {
		var flag = 0
		select {
		case <-ticker:
			flag += rand.Intn(1000)
			fmt.Printf("**********task<%d>**********\n", flag)
		default:
			//time.Sleep(1)
			runtime.Gosched()
		}

	}
}

// fake competitor
func competitor() {
	for {
		fmt.Printf("competitor running...\n")
		// goroutine 是抢占式的，必须显式地让出CPU资源去处理其它任务

		time.Sleep(time.Millisecond * 1000)
		runtime.Gosched()
	}
}

// runner
func run() {
	for {}
	time.Sleep(time.Second * 20)
}

func main() {
	ticker := time.Tick(time.Second * 3)

	go tack(ticker)

	go competitor()
	run()
}
