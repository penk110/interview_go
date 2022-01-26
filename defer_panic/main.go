package main

import "fmt"

func NotRecover() {
	defer func() {
		//e := recover()
		//fmt.Printf("捕捉到了一个bug: %v\n", e)
		fmt.Printf("打印前\n")
	}()
	defer func() {
		fmt.Printf("打印中\n")
	}()
	defer func() {
		fmt.Printf("打印后\n")
	}()

	panic("假设写了个bug")
}

func main() {
	NotRecover()
}

/*
执行结果：
打印后
打印中
打印前
panic: 假设写了个bug

goroutine 1 [running]:
main.NotRecover()
	/Users/zhang/gopath/src/penk110/interview_go/defer_panic/main.go:16 +0xca
main.main()
	/Users/zhang/gopath/src/penk110/interview_go/defer_panic/main.go:23 +0x20
exit status 2

考点：
	defer 执行流程
解答：
	defer为后进先出，协程遇到panic时遍历本协程链表并执行defer，在执行defer过程中遇到recover则停止panic，返回recover处继续往下执行；
	如果执行defer链表过程中没有recover，则遍历完defer链表之后，向stderr抛出panic信息

分析前提是：
	Go 的函数返回值是通过堆栈返回的

*/
