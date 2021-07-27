package sort


import "fmt"

// 最小堆
type minheap struct {
	arr []int
}

// 初始化heap
func newMinHeap(arr []int) *minheap {
	minheap := &minheap{
		arr: arr,
	}
	return minheap
}

// 左孩子节点的位置下标
func (m *minheap) leftchildIndex(index int) int {
	return 2*index + 1
}

// 右孩子节点的位置下标
func (m *minheap) rightchildIndex(index int) int {
	return 2*index + 2
}

// 将左右孩子节点最小的元素和父节点交换位置
func (m *minheap) swap(first, second int) {
	temp := m.arr[first]
	m.arr[first] = m.arr[second]
	m.arr[second] = temp
}

func (m *minheap) leaf(index int, size int) bool {
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}

// 调整heap
// 根据指定的节点位置和堆大小，调整heap
func (m *minheap) downHeapify(current int, size int) {
	if m.leaf(current, size) {
		return
	}
	smallest := current
	leftChildIndex := m.leftchildIndex(current)
	rightRightIndex := m.rightchildIndex(current)
	if leftChildIndex < size && m.arr[leftChildIndex] < m.arr[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < size && m.arr[rightRightIndex] < m.arr[smallest] {
		smallest = rightRightIndex
	}
	// 区左右孩子节点中最小的，如果最小的在子节点的话
	if smallest != current {
		// 先交换元素，然后递归调整最小节点所在的位置，直至叶子节点为止
		m.swap(current, smallest)
		m.downHeapify(smallest, size)
	}
	return
}

// 构建heap
func (m *minheap) buildMinHeap(size int) {
	// 从第一个不是叶子节点的位置开始调整heap
	// 直到根节点，调整才算结束； 根节点为最小元素
	for index := (size / 2) - 1; index >= 0; index-- {
		m.downHeapify(index, size)
	}
}

// 排序的过程：
// 不断将跟节点（最小元素）与 堆末尾元素交互
// 每交换一次，重新调整heap， heap的大小减1
// 交换 size - 1 次之后， 就算排好了序
func (m *minheap) sort(size int) {
	m.buildMinHeap(size)
	for i := size - 1; i > 0; i-- {
		// Move current root to end
		m.swap(0, i)
		m.downHeapify(0, i)
	}
}

// 打印heap
func (m *minheap) print() {
	for _, val := range m.arr {
		fmt.Println(val)
	}
}

// inputArray := []int{6, 5, 3, 7, 2, 8, -1}
// minHeap := newMinHeap(inputArray)
// minHeap.sort(len(inputArray))
// minHeap.print()
// fmt.Scanln()