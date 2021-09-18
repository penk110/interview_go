package main

import (
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func main() {
	/*
		无缓冲通道：
			必须需要有消费端，否则会导致死锁
		goroutine 泄漏
	*/

	var (
		ch1    = make(chan struct{}, 0)
		ch2    = make(chan struct{}, 0)
		ch3    = make(chan struct{}, 0)
		doneCh = make(chan struct{}, 0)
	)

	wg.Add(1)
	go p1(ch3, ch1, doneCh)
	go p2(ch1, ch2, doneCh)
	go p3(ch2, ch3, doneCh)

	// 触发条件
	ch3 <- struct{}{}
	wg.Wait()

	time.Sleep(time.Second * 3)
}

func p1(ch3 chan struct{}, ch1 chan struct{}, done chan struct{}) {
	var i int
	for {
		select {
		case <-ch3: // 无缓冲通道阻塞接收
			i++
			if i > 3 {
				wg.Done()
				println("p1 done")
				close(done)
				return
			}
		}
		println("p1")
		ch1 <- struct{}{}
	}
}

func p2(ch1 chan struct{}, ch2 chan struct{}, done chan struct{}) {
	for {
		select {
		case <-ch1:
			println("p2")
			ch2 <- struct{}{}
		case <-done:
			println("p2 done")
			return
		}
	}
}

func p3(ch2 chan struct{}, ch3 chan struct{}, done chan struct{}) {
	for {
		select {
		case <-ch2:
			println("p3")
			ch3 <- struct{}{}
		case <-done:
			println("p3 done")
			return
		}
	}
}
