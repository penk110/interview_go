package main

import "fmt"

func main() {
	var nilCh chan int

	// fatal error: all goroutines are asleep - deadlock!
	// nilCh <- 10

	// fatal error: all goroutines are asleep - deadlock!
	var rev int
	rev = <-nilCh

	fmt.Println(rev, nilCh)
}

/*
结论：
对nil channel读写会造成永久堵塞
*/
