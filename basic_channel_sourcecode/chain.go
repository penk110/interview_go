package main

import (
	"fmt"
	"time"
)

func main() {
	// 无缓冲通道：向通道写入数据的前提是必须存在另一个携程读取通道，否则进入死锁
	var c1 = make(chan int)
	var exitCh = make(chan struct{})
	//c1 <- 1
	//// 读在之后，导致死锁
	//// fatal error: all goroutines are asleep - deadlock!
	//<-c1

	go func() {
		c1 <- 1
	}()
	go func() {
		ticker := time.NewTicker(time.Second * 2)
		// 阻塞等待定时器触发
		select {
		case <-ticker.C:
			break
		}
		fmt.Println("print after 2s: ", <-c1)
		exitCh <- struct{}{}
	}()

	fmt.Println("exit after 2s: ", <-exitCh)
}
