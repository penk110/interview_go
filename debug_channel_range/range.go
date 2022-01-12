package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int

	ch = make(chan int, 10)
	ch <- 19
	ch <- 18
	ch <- 17
	ch <- 16

	//for c := range ch {
	//	fmt.Println(c)
	//}
	//close(ch)

	fmt.Println(1 << 10)
	// 0011 左移两位 -> 1100
	// 1100 || 0011 -> 1111
	fmt.Printf("%b %v\n", 3<<2|3, 3<<2|3)
	time.Sleep(time.Second * 3)
}
