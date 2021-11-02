package main

import "fmt"

func main() {

	// 写已经关闭的通道导致程序panic
	var c1 = make(chan int)
	close(c1)
	// panic: send on closed channel
	//c1 <- 1

	// 读已经关闭的通道，得到零值panic
	// TODO: 多次关闭将导致程序
	var c2 = make(chan int, 2)
	c2 <- 1
	c2 <- 2
	close(c2)
	for v := range c2 {
		fmt.Println(v)
	}
	// 得到零值
	fmt.Println(<-c2)

}
