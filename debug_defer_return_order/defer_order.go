package main

import (
	"fmt"
)

func p(i *int) int {
	defer func() {
		*i = 19	// 指针操作会改变传进来的变量的实际值 19
	}()

	// 这里是取值返回，所以是取出刚传进来的值 10
	return *i
}

func main() {
	i := 10
	j := p(&i)
	// 19 10
	fmt.Println(i, j)
}
