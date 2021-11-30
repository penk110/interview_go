package main

import (
	"fmt"
	"sync"
	"time"
)

func iteration() {
	var (
		m     map[int]int
		mux   *sync.RWMutex
		timer *time.Timer
		v     int
	)
	m = make(map[int]int)
	mux = &sync.RWMutex{}
	m[1] = 1
	m[2] = 2
	m[3] = 3

	go func(mux *sync.RWMutex, m map[int]int) {
		var (
			timer *time.Timer
			count int
		)
		timer = time.NewTimer(time.Nanosecond * 1)
		for {
			<-timer.C
			mux.Lock()
			m[count] = count
			count++
			mux.Unlock()

			timer.Reset(time.Nanosecond * 1)
		}
	}(mux, m)

	timer = time.NewTimer(time.Millisecond)
	for {
		<-timer.C
		// 遍历不加锁，导致和上面的写操作冲突
		// fatal error: concurrent map iteration and map write
		for k, _ := range m {
			mux.RLock()
			v = m[k]
			fmt.Println(v)
			mux.RUnlock()
		}

		timer.Reset(time.Millisecond)
	}
}

func iterationFix() {
	var (
		m     map[int]int
		mux   *sync.RWMutex
		timer *time.Timer
		v     int
	)
	m = make(map[int]int)
	mux = &sync.RWMutex{}
	go func(mux *sync.RWMutex, m map[int]int) {
		var (
			timer *time.Timer
			count int
		)
		timer = time.NewTimer(time.Nanosecond * 1)
		for {
			<-timer.C
			mux.Lock()
			m[count] = count
			count++
			mux.Unlock()

			timer.Reset(time.Nanosecond * 1)
		}
	}(mux, m)

	timer = time.NewTimer(time.Millisecond)
	for {
		<-timer.C
		// 遍历不加锁，导致和上面的写操作冲突
		// fatal error: concurrent map iteration and map write
		mux.RLock()
		for k, _ := range m {
			v = m[k]
			fmt.Println(v)
		}
		mux.RUnlock()
		timer.Reset(time.Millisecond)
	}
}

func main() {
	// 并发异常
	// iteration()

	// 解决
	iterationFix()
}
