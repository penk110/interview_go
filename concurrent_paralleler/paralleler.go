package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	signal := make(chan struct{})
	parallelCount := 100
	wg := &sync.WaitGroup{}

	for i := 1; i <= parallelCount; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, single chan struct{}, i int) {
			<-single
			fmt.Printf("[Paralleler] index: %v  time: %v\n", i, time.Now().Format("2006-01-02 15:04:05.0000"))
			wg.Done()
		}(wg, signal, i)
	}
	fmt.Println("waiting signal")
	fmt.Println("sleep 2s ...")
	go func() {
		i := 10
		for ; i > 0; {
			fmt.Println("waiting ...")
			time.Sleep(time.Microsecond * 5000)
			i--
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println(" trigger signal ...")
	close(signal)
	wg.Wait()
}
