package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

/*
	select 多路复用
	1.与switch不同，每个case语句必须对应通道的读写操作，
	2.select语句会陷入阻塞，直到一个或者多个通道能够正常读写

	特性：
	1.随机读写(随机选择可读写的通道)
	2.堵塞与控制，没有任何通道准备好，则进入阻塞
		常用于：1)方法返回; 2)
	3.default 存在该分支时如其它通道均阻塞则执行default分支
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
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			c <- strconv.FormatInt(time.Now().Unix(), 10)
		case <-ctx.Done():
			return
		default:
			fmt.Printf("default case\n")
			time.Sleep(time.Second)
		}

	}
	/*
	default case
	default case
	default case
	worker id: 123 rcv: 1635786067
	default case
	default case
	worker id: 123 rcv: 1635786069
	default case
	default case
	worker id: 123 rcv: 1635786071
	default case
	default case
	worker id: 123 rcv: 1635786073
	default case
	*/
}
