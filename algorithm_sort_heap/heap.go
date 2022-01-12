package main

import "fmt"

/*
关键过程为建堆、堆顶与无序区尾部节点交换

# Runtime:
- Average: O(n log n)
- Worst: O(n log n)
- Best: O(n log n)

解析:
	1.已知父子节点的索引为i的数，即可计算出左右子节点的索引
		父节点索引: (i-1)/2
		左节点索引: 2*i + 1
		右边节点索引: 2*i + 2
	2.父节点和左右子节点大小关系 --- 用于构建堆的判断依据
		大根堆：arr(i)>arr(2*i+1) && arr(i)>arr(2*i+2)
		小根堆：arr(i)<arr(2*i+1) && arr(i)<arr(2*i+2)

步骤：
	1.构建堆(升序：大根堆；降序：小根堆)
		数组映射成大堆 <==> 完全二叉树从上到下，从左到右遍历得到的元素组合层的列表

	2.每次将堆顶移除，循环此过程即可得到有序列表

参考：
	https://blog.csdn.net/m0_46803965/article/details/105612093
	https://blog.csdn.net/weixin_42654444/article/details/83513488?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.channel_param&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.channel_param


堆排序思想：
	将待排序序列构造成一个大顶堆，此时，整个序列的最大值就是堆顶的根节点。将其与末尾元素进行交换，此时末尾就为最大值
	然后将剩余n-1个元素重新构造成一个堆，这样会得到n个元素的次小值。如此反复执行，便能得到一个有序序列了
*/

// Heap is heap struct
type Heap struct {
	Data      []int
	Count     int
	HeapType  string
	HeapCount int
}

// NewHeap is create new heap func and return heap object
func NewHeap(data []int, capacity int, t string) *Heap {
	return &Heap{
		Data:     data,
		Count:    capacity,
		HeapType: t,
	}
}

// BuildMaxHeap is build max heap and return heap
func (h *Heap) BuildMaxHeap() {
	// 构建大堆，对 倒数第二层最左节点 开始进行
	// 层数计算公式： level = Count/2 - 1
	for i := h.Count/2; i >= 0; i-- {
		h.Heapify(i, h.Count)
	}
}

// BuildMinHeap is build min heap and return heap
func (h *Heap) BuildMinHeap() {
	// 构建小堆，对第二层节点开始进行
	for i := 0; i <= h.Count/2; i++ {
		h.Heapify(i, h.Count)
	}
}

// Heapify is build a heap
func (h *Heap) Heapify(i, length int) {
	// 记录递归层数
	h.HeapCount++

	// 明确跳出递归条件
	// fmt.Printf("heapiy level: %v\n", i)
	//if i > length {
	//	return
	//}

	// 左右子节点索引
	left := 2*i + 1
	right := 2*i + 2
	// 默认父子节点最大
	largest := i

	// 判断获取最大节点的索引
	if right < length && h.Data[right] > h.Data[largest] {
		largest = right
	}
	if left < length && h.Data[left] > h.Data[largest] {
		largest = left
	}
	if largest != i {
		h.Swap(i, largest)

		// 设置递归进入条件可减少递归次数
		if largest*2+1 < length {
			h.Heapify(largest, length)
		}
	}
	return
}

// Swap is swap slice elm function
func (h *Heap) Swap(i, j int) {
	// 交换两个元素
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

// Sort is sort heap func and return sorted slice
func (h *Heap) Sort() []int {
	if h.HeapType == "max" {
		h.BuildMaxHeap()
		fmt.Printf("build max heap after: %v\nheapify count: %v\n", h.Data, h.HeapCount)

		// 排序过程：每次把最后一个节点（最后一层的最右边的节点）和堆顶（始终是最大的，第一个节点）进行交换
		// 最终得到一个升序的序列
		for i := h.Count - 1; i >= 0; i-- {
			h.Swap(0, i)
			h.Count -= 1
			h.Heapify(0, h.Count)
		}
	} else if h.HeapType == "min" {
		// TODO: 小堆构建为完成
		h.BuildMinHeap()
		fmt.Printf("build min heap after: %v\n", h.Data)
		for i := 0; i <= h.Count-1; i++ {
			h.Swap(0, i)
			h.Count -= 1
			h.Heapify(0, h.Count)
		}
	}

	return h.Data
}

func main() {
	heapType := "max"

	data := []int{
		12, 123, 44, 22, 54, 1, 2, 4, 56, 3,
	}
	heap := NewHeap(data, len(data), heapType)

	fmt.Printf("before: %v\n", data)
	data = heap.Sort()
	// sortData = [1 2 3 4 12 22 44 54 56 123]
	fmt.Printf("after: %v\n", data)
}
