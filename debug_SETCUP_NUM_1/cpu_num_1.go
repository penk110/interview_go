package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Microsecond * 300)
			println(i)
		}(i)
	}

	time.Sleep(time.Second * 3)

}

/*
先输出 9
原因是：
1.next queue和local queue两个队列中没有G，只能从全局global中获取
2.从global queue中获取G时，先将最后一个G取出来放到next queue 中，其余在取 local queue 长度的一半个数个G？后面确定下个数

9
0
1
2
3
4
5
6
7
8
*/
