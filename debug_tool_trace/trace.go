package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func mockSendToServer(url string) {
	fmt.Printf("server url: %s\n", url)
}

func main() {
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	urls := []string{"0.0.0.0:8081", "0.0.0.0:8082", "0.0.0.0:8083"}

	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mockSendToServer(url)
		}()
	}
	wg.Wait()
}

/*
运行之后可获取到trace.out文件

通过tool工具进行分析
	go tool trace trace.out

Goroutine Name:	main.main.func1
Number of Goroutines:	3
Execution Time:	60.02% of total program execution time
Network Wait Time:	graph(download)
Sync Block Time:	graph(download)
Blocking Syscall Time:	graph(download)
Scheduler Wait Time:	graph(download)
*/