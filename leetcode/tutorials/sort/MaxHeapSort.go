package sort

// MinHeapSort 使用最小堆排序算法对切片 a 进行排序
// 参数:
//     a []int: 待排序的整数切片
func MinHeapSort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		adjustHeap(a, 0, i, minHeapOpt())
	}
}
func MaxHeapSort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		adjustHeap(a, 0, i, maxHeapOpt())
	}
}
func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2 ; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}
func maxHeapify(a []int, index int, heapSize int) {
	// 从需要调整的index位置节点开始和左右节点比较
	// 大顶堆： 找到最大的元素
	// 小顶堆： 找到最小的元素
	// 如果最小或者最大的元素就是在index节点，那就不用调整
	// 如果需要调整，那么需要递归往下看，该调整的元素是否都比左右节点小或者大（从根节点调整下来的那个元素可能比左右节点的元素小或者大）
	left, right, largest := 2 * index + 1, 2*index + 2, index
	if left < heapSize && a[left] > a[index]  {
		largest = left
	}
	if right < heapSize && a[right] > a[largest] {
		largest = right
	}
	if largest != index {
		// 交换
		a[index], a[largest] = a[largest], a[index]
		maxHeapify(a, largest, heapSize)
	}
}



// AdjustOpt 是一个函数类型，用于定义堆调整操作
type AdjustOpt func([]int, int, int) int

//  大顶堆：
// maxHeapOpt 返回一个 AdjustOpt 类型，用于调整最大堆
//
// 参数：
// 无
//
// 返回值：
// AdjustOpt：一个用于调整最大堆的 AdjustOpt 类型
func  maxHeapOpt() AdjustOpt {
	return AdjustOpt(func(a []int, index, heapSize int) int {
		cur := index
		left, right := 2 * index + 1, 2*index + 2
		if left < heapSize && a[left] > a[cur]  {
			cur = left
		}
		if right < heapSize && a[right] > a[cur] {
			cur = right
		}
		return cur
	})
}
// 小顶堆
// minHeapOpt 返回一个用于最小堆优化的调整函数。
// 该函数返回一个 AdjustOpt 类型的匿名函数，该匿名函数接受一个整数切片 a，
// 一个索引 index，以及堆的大小 heapSize 作为参数，返回一个整数类型的索引值。
//
// 参数说明：
// - a []int: 待调整的整数切片
// - index int: 当前需要调整的节点索引
// - heapSize int: 堆的大小
//
// 返回值：
// - int: 调整后的节点索引
func minHeapOpt() AdjustOpt {
	return func(a []int, index, heapSize int) int {
		cur := index
		left, right := 2 * index + 1, 2*index + 2
		if left < heapSize && a[left] < a[cur]  {
			cur = left
		}
		if right < heapSize && a[right] < a[cur] {
			cur = right
		}
		return cur
	}
}


// NewMinHeap 函数创建一个最小堆
//
// 参数:
// a: 一个整数切片，用于构建堆
// heapSize: 堆的大小
func NewMinHeap(a []int, heapSize int) {
	buildHeap(a, heapSize, minHeapOpt())
}

// NewMaxHeap 创建一个最大堆
// 参数 a 表示需要构建最大堆的切片
// 参数 heapSize 表示堆的大小
func NewMaxHeap(a []int, heapSize int) {
	buildHeap(a, heapSize, maxHeapOpt())
}


// buildHeap 根据给定选项 opt 构建一个大小为 heapSize 的堆
// 参数：
//   - a: 待构建堆的数组
//   - heapSize: 堆的大小
//   - opt: 堆调整选项
func buildHeap(a []int, heapSize int, opt AdjustOpt) {
	for i := heapSize / 2  ; i >= 0; i-- {
		adjustHeap(a, i, heapSize, opt)
	}
}

// adjustHeap 对给定的切片 a 进行堆调整
// a 是一个整数切片，index 是要调整的堆节点的索引
// heapSize 是堆的大小，maxOrMinCmp 是一个函数，用于比较两个元素的大小，以确定堆是最大堆还是最小堆
func adjustHeap(a []int, index int, heapSize int, maxOrMinCmp AdjustOpt) {
	if index >= heapSize || index < 0 {
		return
	}
	if current := maxOrMinCmp(a, index, heapSize); current != index {
		// 交换
		a[index], a[current] = a[current], a[index]
		adjustHeap(a, current, heapSize, maxOrMinCmp)
	}
}
