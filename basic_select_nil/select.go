package main

import (
	"fmt"
)

/*
	select 多路复用 结合 nil
*/

func main() {
	var c1 = make(chan int)
	var c2 = make(chan int)
	go func() {
		for i := 1; i <= 2; i++ {
			select {
			case c1 <- i:
				c1 = nil // 使得第二次循环不会走到 case c1 <- i 而是顺序执行 case c2 <- i
			case c2 <- i:
				c2 = nil
			}
		}
	}()
	fmt.Println("c1: ", <-c1)
	fmt.Println("c2: ", <-c2)
}
