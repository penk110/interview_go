package main

import "fmt"

type slice []int

func NewSlice() slice {
	return make([]int, 0)
}

func (a *slice) add(i int) *slice {
	*a = append(*a, i)
	fmt.Println(i)
	return a
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	s := NewSlice()
	// 函数堆栈
	defer func() {
		// 作为整体包在匿名函数中会 延迟执行
		s.add(1).add(2)
	}()

	s.add(3)

	r := increaseB()
	fmt.Println(r)
}
