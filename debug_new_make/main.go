package main

import (
	"fmt"
	"reflect"
)

// wrong define
func wrongDefine() {
	var i *int

	// 不使用 new 初始化进行内存分配，报错：
	// panic: runtime error: invalid memory address or nil pointer dereference
	i = new(int)
	*i = 10

	fmt.Println(*i)
}

func main() {
	// 错误定义
	wrongDefine()

	var i int
	fmt.Printf("%b %v %b\n", i, reflect.TypeOf(i), 10)

	// make(v, length, cap)
	// 不指定cap，则在make初始化时会默认设置等于length
	s1 := make([]int, 1)
	s2 := make([]int, 1, 2)
	s3 := make([]int, 0, 0)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
}
