package main

import (
	"fmt"
)

func selfRecover() interface{} {
	fmt.Printf("self recover\n")
	return recover()
}

func main() {
	defer func() {
		r := selfRecover()
		fmt.Printf("自定义无效的recover: <%v>！！！\n", r)
	}()

	// 正常捕捉和处理异常
	panic("defer间接调用recover无效！")
}

/*
/private/var/folders/wp/0kxj_ktn6xqg0qrrrty0jcr80000gn/T/___go_build_zyphub_interview_go_debug_error_recover2
self recover
自定义无效的recover: <<nil>>！！！
panic: defer间接调用recover无效！

goroutine 1 [running]:
main.main()
        /Users/zhang/gopath/src/zyphub/interview_go/debug_error_recover2/error.go:19 +0x5f
*/
