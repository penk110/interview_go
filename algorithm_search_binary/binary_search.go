package main

import (
	"fmt"
	"sort"
)

/*
	二分查找
	1.确保列表是有序的


*/

func BinarySearch(v int, l []int) int {
	n := len(l)
	if n == 0 {
		return -1
	}

	return bs(v, 0, n-1, l)
}

func bs(v int, lIndex int, hIndex int, n []int) int {
	if lIndex > hIndex {
		return -1
	}
	mid := (lIndex + hIndex) >> 1
	m := n[mid]
	if m == v {
		return mid
	} else if m > v {
		// 目标值 小于中间值
		return bs(v, lIndex, mid-1, n)
	} else {
		// 目标值 大于 中间值
		return bs(v, mid+1, hIndex, n)
	}
}

func main() {
	lIndex, hIndex := 1, 4
	// 	右移运算符">>"是双目运算符
	//	右移n位就是除以2的n次方
	//	其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数
	mid := (lIndex + hIndex) >> 1

	fmt.Println(mid)

	// 返回列表索引
	n := []int{1, 5, 6, 9, 12, 33}
	sort.Ints(n)
	//fmt.Println(BinarySearch(3, n))
	fmt.Println(BinarySearch(9, n))
}
