package main

import "fmt"

func a() (r int) {
	defer func() {
		// defer 在 return之前执行，defer内是指针传递，所以会直接修改r的值 r=1，最终返回 1
		r++
	}()
	// r 返回的r被赋值为0
	return 0
}

func b() (r int) {
	t := 5
	// r=5
	r = t
	defer func() {
		// 在 return 之前，执行 defer 函数
		// defer 函数没有对返回值 r 进行修改，只是修改了变量 t
		t = t + 5
	}()
	return
}

func c() (r int) {
	// 赋值
	r = 1
	// 这里修改的 r 是函数形参的值
	// 值拷贝，不影响实参值
	func(r int) {
		r = r + 5
	}(r)
	return
}


func main() {
	fmt.Println("a: ", a())
	fmt.Println("b: ", b())
	fmt.Println("c: ", c())
}
