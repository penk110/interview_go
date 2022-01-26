package main

import (
	"flag"
	"log"
	"os"
	"path"
	"runtime/pprof"
	"time"
)

func task() {
	var ch chan int
	for {
		select {
		case value := <-ch:
			log.Printf("recive value<%v> from channel\n", value)
		default:
			// Gosched yields the processor, allowing other goroutines to run. It does not
			// suspend the current goroutine, so execution resumes automatically.
			//runtime.Gosched()
		}
	}
}

func main() {
	var CPUPprof bool
	var MEMPprof bool

	flag.BoolVar(&CPUPprof, "cpu", false, "turn cpu pprof on!")
	flag.BoolVar(&MEMPprof, "mem", false, "turn mem pprof on!")
	// flag parse
	flag.Parse()

	pwd, _ := os.Getwd()
	if CPUPprof {
		file, err := os.Create(path.Join(pwd, "cpu.pprof"))
		if err != nil {
			log.Panicf("create cpu pprof file failed, err: %v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	// start task
	for i := 0; i < 10; i++ {
		go task()
	}

	time.Sleep(time.Second * 10)
	if MEMPprof {
		file, err := os.Create(path.Join(pwd, "mem.pprof"))
		if err != nil {
			log.Panicf("create mem pprof file failed, err: %v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		defer file.Close()
	}

}

/*
一、命令行交互
go tool pprof cpu.pprof

查看top，输入前三的：top3
	(pprof) top
	Showing nodes accounting for 81.16s, 100% of 81.18s total
	Dropped 7 nodes (cum <= 0.41s)
		  flat  flat%   sum%        cum   cum%
		33.29s 41.01% 41.01%     64.80s 79.82%  runtime.selectnbrecv
		25.21s 31.05% 72.06%     27.24s 33.56%  runtime.chanrecv
		16.36s 20.15% 92.21%     81.16s   100%  main.task
		 6.30s  7.76%   100%      6.30s  7.76%  runtime.newstack
	(pprof) list task
	Total: 1.35mins
	ROUTINE ======================== main.task in /Users/zhang/gopath/src/penk110/interview_go/anaylisis_pprof/temp.go
		16.36s   1.35mins (flat, cum)   100% of Total
			 .          .     12:
			 .          .     13:func task() {
			 .          .     14:   var ch chan int
			 .          .     15:   for {
			 .          .     16:           select {
		16.36s   1.35mins     17:           case value := <-ch:
			 .          .     18:                   log.Printf("recive value<%v> from channel\n", value)
			 .          .     19:           default:
			 .          .     20:                   // Gosched yields the processor, allowing other goroutines to run. It does not
			 .          .     21:                   // suspend the current goroutine, so execution resumes automatically.
			 .          .     22:                   //runtime.Gosched()



解释：
	flat：当前函数占用CPU的耗时
	flat：:当前函数占用CPU的耗时百分比
	sun%：函数占用CPU的耗时累计百分比
	cum：当前函数加上调用当前函数的函数占用CPU的总耗时
	cum%：当前函数加上调用当前函数的函数占用CPU的总耗时百分比
	最后一列：函数名称

二、图形化分析
1.将pprof文件转换成pdf文件：
	go tool pprof --pdf cpu.pprof > cpu.pdf
2.生成CSV热力图
	go-torch anaylisis_pprof/cpu.pprof
*/
