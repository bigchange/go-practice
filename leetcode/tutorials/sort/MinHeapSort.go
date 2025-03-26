package sort

import "fmt"

// 最小堆
type minheap struct {
	arr []int
}

// 排序：
// 1. 首先建好堆，使其符合根节点为最小元素的堆
// 2. 其次，对堆进行排序，也就是每次将（堆顶）最小元素放到数组末尾，然后调整堆顶元素使其符合堆的性质，
// 3. 重复上述过程直到调整了size - 1 次（每次调整heap的大小减1），也就将所有元素都排好了序
// 排序简要过程：
// 不断将建好的堆的根节点（root节点，即为最小元素）与 堆末尾元素交互
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

// 初始化heap
// 建堆： O(n)
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
// 找到当前节点和左右孩子节点中最小元素的位置
// 1. 如果当前节点就是最小的，则不需要调整
// 2. 如果不是，则将最小元素和当前节点交换位置， 然后调整最小元素所在位置元素（该元素是已经swap后的值，不是最小值了）
// 递归调整最小节点所在的位置，直至叶子节点为止	
// 这样就能保证根节点都是比子节点小的
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
	// 取左右孩子节点中最小的，如果最小的在子节点的话
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
