package main

import "fmt"

func f1() int {
	i := 3
	defer func() {
		i++
		fmt.Printf("inside defer: id<%v>, value<%v>\n", &i, i)
		// inside defer: id<0xc000098008>, value<4>
	}()
	fmt.Printf("outside defer: id<%v>, value<%v>\n", &i, i)
	// outside defer: id<0xc000098008>, value<3>
	return i
}

func f2() (i int) {
	defer func() {
		i++
	}()
	return 0
}

func c() (i int) {
	defer func() { i++ }()

	// 先把i赋值为1，defer中又对i进行自增1，所以最终返回2
	return 1
}

func main() {
	//
	fmt.Println(f1())

	//
	fmt.Println(f2())

	fmt.Println(c())
}

/*
核心点：
	延迟函数延迟执行，


*/
