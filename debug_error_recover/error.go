package main

import (
	"fmt"
)

func main() {
	user := make(map[string]interface{})
	user["name"] = "golang"
	user["age"] = 18
	fmt.Printf("user: %#v\n", user)

	// ok 返回map中是否存在这个key，存在则为 true, 否则为 false
	if value, ok := user["phone"]; ok {
		fmt.Printf("user phone: %v\n", value)
	} else {
		fmt.Printf("involid key: %v, error: %v\n", "phone", ok)
	}

	if e := recover(); e == nil {
		fmt.Println("直接调用recover执行的panic发生之前的正常流程，不会捕捉到异常")
	}

	defer func() {
		r := recover()
		fmt.Printf("panic 之前进行defer堆栈、注册，panic之后会正常捕捉到异常: <%v>！！！\n", r)
	}()

	// 正常捕捉和处理异常
	panic("我决定在recover语句之后才panic，所以上面那个panic休想捕捉到我！")

	//defer func() {
	//	r := recover()
	//	fmt.Printf("panic 之后无法正常进行defer堆栈、注册，无法正常捕捉到异常<%v>！！！\n", r)
	//}()

	// 这里肯定不会被执行，因为上面已经被panic了
	if r := recover(); r != nil {
		fmt.Printf("很明显我是在被panic之后才执行的，没有被注册到函数中，怎么可能捕捉到panic？？？ %v\n", r)
	}
}

/*
查看翻译成汇编的伪代码：
	go tool compile -S error.go
*/
