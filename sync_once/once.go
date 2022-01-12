package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singletion struct {
}

var (
	instance    *singletion
	initialized uint32
	mu          sync.RWMutex
)

// Instance 双重锁，上锁前先判断是否被实例化即 是否为nil
func Instance() *singletion {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if instance != nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singletion{}
	}
	return instance
}

// 上面的方法实际上是从 sync.Once() 提取出来的，可直接使用原始的Once实现

var once sync.Once

// InstanceOnce insure create single instance
func InstanceOnce() *singletion {
	once.Do(func() {
		if instance == nil {
			instance = &singletion{}
		}
	})

	return instance
}

func main() {
	var m sync.Map
	m = sync.Map{}
	m.Store("addr", "127.0.0.1")

	addr, ok := m.Load("addr")
	if !ok {
		fmt.Println("key does not exists!")
	}
	fmt.Printf("sync map: %v\n", addr)
}
