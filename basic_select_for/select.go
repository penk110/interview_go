package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
	select 多路复用 结合 for
*/

func Worker(wid string, c chan string) {
	for v := range c {
		fmt.Printf("worker id: %s rcv: %v\n", wid, v)
	}
}

func Master(id string) chan string {
	var c = make(chan string)
	go Worker(id, c)

	return c
}

func main() {
	c := Master("123")
	ticker := time.NewTicker(time.Second * 2)
	//for {
	//	select {
	//	case <-ticker.C: // ticker定时往管道内塞数据，即实现定时
	//		c <- strconv.FormatInt(time.Now().Unix(), 10)
	//	case <-time.After(time.Second * 3): // 每次都会被重置，无意义，下次select还是需要等待3s后，如果当前case存在已准备好的通道，则当前case永远得不到执行
	//		fmt.Println("after 10s")
	//		return
	//	}
	//}

	for {
		select {
		case <-ticker.C:
			c <- strconv.FormatInt(time.Now().Unix(), 10)
		case <-time.After(time.Millisecond * 3): // 有机会循环到并执行
			fmt.Println("after time.Millisecond*3")
		}
	}
}
