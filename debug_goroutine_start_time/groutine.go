package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func mockSendToServer(url string) {
	fmt.Printf("server url: %s\n", url)
}

func main() {
	//debug.SetMaxThreads(5)

	urls := []string{"0.0.0.0:8081", "0.0.0.0:8082", "0.0.0.0:8083"}
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		u := url
		go func() {
			defer wg.Done()
			mockSendToServer(u)
		}()
	}

	// 查看当前程序运行了多少个goroutine
	fmt.Printf("runtime.NumGoroutine: %v\n", runtime.NumGoroutine())
	wg.Wait()

	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			//fmt.Printf("goroutine<%d> is running...\n", id)
			time.Sleep(time.Microsecond * 100)

		}(i)
	}

	fmt.Printf("runtime.NumGoroutine: %v\n", runtime.NumGoroutine())
	wg.Wait()
}

/*
逻辑：
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mockSendToServer(url)
		}()
	}
输出：
	server url: 0.0.0.0:8083
	server url: 0.0.0.0:8083
	server url: 0.0.0.0:8083
分析：
	url为局部变量，已上逻辑实现了闭包，即goroutine的func是对url进行引用，启动和执行时间刚好是循环到最后一个url

解决：
	将URL以参数形式传进goroutine中
		for _, url := range urls {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				mockSendToServer(u)
			}(url)
		}
	或者：
		for _, url := range urls {
			wg.Add(1)
			u := url
			go func() {
				defer wg.Done()
				mockSendToServer(u)
			}()
		}

需要注意的是：
	最终urls元素输出的顺序是不确定的，这和goroutine启动和执行时间有关

*/
