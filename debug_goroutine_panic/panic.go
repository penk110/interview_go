package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	// 内部不做异常处理会导致主线程panic退出
	defer func() {
		_ = recover()
	}()
	time.Sleep(time.Second * 1)
	if i%10 == 0 {
		panic("task<10> panic")
	}
	fmt.Printf("task<%v> is running...\n", i)
	return
}

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)

	// 这里捕捉异常是没有意义的，因为内部已经执行了 panic
	defer func() {
		r := recover()
		fmt.Println(r)
	}()

	var wg sync.WaitGroup
	wg.Add(15)
	for i := 1; i <= 15; i++ {
		task(&wg, i)
	}
	wg.Wait()
	fmt.Println("all task done!")
}
