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


	select一轮循环:
		一轮循环，对所有scase循环(caseNil caseRecv caseSend caseDefault)，找到已就绪的通道(caseNil会被忽略)，caseDefault默认执行
	select二轮循环:
		一轮循环未退出，对当前协程需要进入休眠状态，并等待select至少有一个通道被唤醒，读取、写入 通道都需创建新的sudog并将其放到指定通道的等待队列中，之后当前协程进入休眠状态

	源码：
		runtime/select.go
		func selectgo(cas0 *scase, order0 *uint16, pc0 *uintptr, nsends, nrecvs int, block bool) (int, bool) {}
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
