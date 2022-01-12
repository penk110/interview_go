package main

import (
	"fmt"
)

// Recursive: 普通递归
func Recursive(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * Recursive(n-1)
}

// RecursiveTail: 尾递归
func RecursiveTail(n, r int64) int64 {
	if n == 0 {
		return r
	}
	return RecursiveTail(n-1, r*n)
}

// 递归实现二分查找
func BinarySearch(a []int, target, l, r int) int {
	if l > r {
		// 出界
		return -1
	}
	middle := (l + r) / 2
	middleNum := a[middle]
	if middleNum == target {
		return middle
	} else if middleNum > target {
		// 中间的数比目标还大，从左边找
		return BinarySearch(a, target, 1, middle-1)
	} else {
		// 中间的数比目标还小，从右边找
		return BinarySearch(a, target, middle+1, r)
	}
}

// 递归实现二分查找
func BinarySearchFor(a []int, target, l, r int) int {
	lTmp := l
	rTmp := r

	for {
		if lTmp > rTmp {
			// 出界
			return -1
		}
		middle := (lTmp + rTmp) / 2
		middleNum := a[middle]
		if middleNum == target {
			return middle
		} else if middleNum > target {
			// 中间的数比目标还大，从左边找
			rTmp = middle - 1
		} else {
			// 中间的数比目标还小，从右边找
			lTmp = middle + 1
		}
	}
}

func main() {
	// 递归式使用了运算符，每次重复的调用都使得运算的链条不断加长，系统不得不使用栈进行数据保存和恢复
	// 可能会存在栈溢出，导致程序奔溃
	fmt.Println(Recursive(10))

	// 规避栈溢出
	fmt.Println(RecursiveTail(10, 1))

	// 二分查找
	tmpArray := []int{1, 20, 30, 153, 813, 1000}
	target1 := 50
	target2 := 30
	fmt.Println(target1, BinarySearch(tmpArray, target1, 0, len(tmpArray)-1))
	fmt.Println(target2, BinarySearch(tmpArray, target2, 0, len(tmpArray)-1))

	fmt.Println(1000, BinarySearchFor(tmpArray, 1000, 0, len(tmpArray)))
	fmt.Println(target2, BinarySearchFor(tmpArray, target2, 0, len(tmpArray)))
}
