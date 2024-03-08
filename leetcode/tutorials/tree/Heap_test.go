package tree

import (
	"container/heap"
	"fmt"
	"testing"
)

/* Driver Code */
func TestHeap(t *testing.T) {
	/* 初始化堆 */
	// 初始化大顶堆
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	/* 元素入堆 */
	// 调用 heap.Interface 的方法，来添加元素
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 5)

	/* 获取堆顶元素 */
	top := maxHeap.Top()
	fmt.Printf("堆顶元素为 %d\n", top)

	/* 堆顶元素出堆 */
	// 调用 heap.Interface 的方法，来移除元素
	heap.Pop(maxHeap) // 5
	heap.Pop(maxHeap) // 4
	heap.Pop(maxHeap) // 3
	heap.Pop(maxHeap) // 2
	heap.Pop(maxHeap) // 1

	/* 获取堆大小 */
	size := len(*maxHeap)
	fmt.Printf("堆元素数量为 %d\n", size)

	/* 判断堆是否为空 */
	isEmpty := len(*maxHeap) == 0
	fmt.Printf("堆是否为空 %t\n", isEmpty)
}
