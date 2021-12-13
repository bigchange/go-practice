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
	if right < heapSize && a[right] > largest {
		largest = right
	}
	if largest != index {
		// 交换
		a[index], a[largest] = a[largest], a[index]
		maxHeapify(a, largest, heapSize)
	}
}

func NewMinHeap(a []int, heapSize int) {
	// 小顶堆
	minHeap  := func(left, right, index int) int {
		cur := index
		if left < heapSize && a[left] < a[index]  {
			cur = left
		}
		if right < heapSize && a[right] < cur {
			cur = right
		}
		return cur
	}
	buildHeap(a, heapSize, minHeap)
}

func NewMaxHeap(a []int, heapSize int) {
	//  大顶堆：
	maxHeap := func(left, right, index int) int {
		cur := index
		if left < heapSize && a[left] > a[index]  {
			cur = left
		}
		if right < heapSize && a[right] > cur {
			cur = right
		}
		return cur
	}
	buildHeap(a, heapSize, maxHeap)
}


func buildHeap(a []int, heapSize int, maxOrMinCmp func(int, int, int) int) {
	for i := heapSize / 2 ; i >= 0; i-- {
		adjustHeap(a, i, heapSize, maxOrMinCmp)
	}
}

func adjustHeap(a []int, index int, heapSize int, maxOrMinCmp func(int, int ,int) int) {
	left, right, largest := 2 * index + 1, 2*index + 2, index
	if current := maxOrMinCmp(index, left, right); current != index {
		// 交换
		a[index], a[largest] = a[largest], a[index]
		adjustHeap(a, largest, heapSize, maxOrMinCmp)
	}
}
