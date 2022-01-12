package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
)

var sum int64

func task() {
	j := rand.Int63n(10000)
	atomic.AddInt64(&sum, j)
	//fmt.Printf("task<%d> is running...\n", j)
}

func main() {
	// 资源释放
	defer ants.Release()

	var wg sync.WaitGroup

	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		sumTask := func() {
			task()
			wg.Done()
		}
		_ = ants.Submit(sumTask)
	}
	fmt.Printf("goroutines num: %d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		task()
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}
