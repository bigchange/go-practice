package sort

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

type AdjustOpt func([]int, int, int) int

//  大顶堆：
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


func NewMinHeap(a []int, heapSize int) {
	buildHeap(a, heapSize, minHeapOpt())
}

func NewMaxHeap(a []int, heapSize int) {
	buildHeap(a, heapSize, maxHeapOpt())
}


func buildHeap(a []int, heapSize int, opt AdjustOpt) {
	for i := heapSize / 2  ; i >= 0; i-- {
		adjustHeap(a, i, heapSize, opt)
	}
}

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
