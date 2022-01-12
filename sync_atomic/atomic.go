package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total int64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		atomic.AddInt64(&total, int64(i))
	}
}

func main() {
	var wg = &sync.WaitGroup{}
	wg.Add(3)
	go worker(wg)
	go worker(wg)
	go worker(wg)
	wg.Wait()

	fmt.Printf("total value: %v\n", total)
}
