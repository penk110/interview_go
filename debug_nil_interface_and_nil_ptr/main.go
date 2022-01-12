package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 定义和初始化，初始化值为 nil，即该指针指向的值为空
	var nilPtr *int = nil
	// 定义和初始化一个空接口，并且指定空接口的值为空
	var b interface{} = nil
	var c interface{}

	fmt.Println(nilPtr == nil, b == nil, nilPtr == b)
	fmt.Println(reflect.TypeOf(nilPtr), reflect.TypeOf(b))
	c = int(10)
	fmt.Println(nilPtr, &nilPtr, b, &b, c, &c)

	//*nilPtr = 1
	//fmt.Println(nilPtr, &nilPtr)
	//panic: runtime error: invalid memory address or nil pointer dereference

	d := 1
	nilPtr = &d
	fmt.Println(nilPtr, &nilPtr, &d)

}
