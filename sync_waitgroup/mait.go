package main

import (
	"fmt"
	"sync"
)

type total struct {
	wait  *sync.WaitGroup
	lock  sync.Mutex
	value int
}

func worker(t *total) {
	defer t.wait.Done()

	for i := 0; i <= 1000; i++ {
		t.lock.Lock()
		t.value += i
		t.lock.Unlock()
	}
}

func main() {
	var t = &total{
		wait:  &sync.WaitGroup{},
		value: 0,
	}
	t.wait.Add(2)
	go worker(t)
	go worker(t)

	t.wait.Wait()
	fmt.Printf("value: %v\n", t.value)
}
