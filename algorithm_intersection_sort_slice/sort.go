package main

import (
	"fmt"
	"sort"
)

/*
	排序+双指针

1.先排序
2.初始化两个列表初始指针
3.遍历、递增初始指针

时间复杂度：O(m*log m + n*log n)
空间复杂度：O(log m+log n)
*/

func intersection(nums1, nums2 []int) []int {
	var (
		res []int
	)
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			// 结果为空或者大于结果列表最大值
			if len(res) == 0 || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return res
}

func main() {
	nums1 := []int{10, 12, 2, 34, 4}
	nums2 := []int{34, 32, 7, 89, 98}
	fmt.Println(intersection(nums1, nums2))
}
