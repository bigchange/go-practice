package tree

// refer: https://www.hello-algo.com/chapter_heap/build_heap/#921
// 「堆Heap」是一种满足特定条件的完全二叉树
// 1. 「大顶堆 Max Heap」，任意节点的值 >= 其子节点的值；
// 2. 「小顶堆 Min Heap」，任意节点的值 <= 其子节点的值；

// 使用场景：
// 1. 堆通常用作实现优先队列，大顶堆相当于元素按从大到小顺序出队的优先队列
// 2. 堆排序：给定一组数据，我们可以用它们建立一个堆，然后依次将所有元素弹出，从而得到一个有序序列。
//    当然，堆排序的实现方法并不需要弹出元素，而是每轮将堆顶元素交换至数组尾部并缩小堆的长度。
// 3. 获取最大的K个元素：这是一个经典的算法问题，同时也是一种典型应用，
//    例如选择热度前 10 的新闻作为微博热搜，选取销量前 10 的商品等。

// Go 语言中可以通过内置数据结构实现 heap.Interface 来构建整数大顶堆
// 实现 heap.Interface 需要同时实现 sort.Interface
type intHeap []any

// Push heap.Interface 的方法，实现推入元素到堆
func (h *intHeap) Push(x any) {
	// Push 和 Pop 使用 pointer receiver 作为参数
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(int))
}

// Pop heap.Interface 的方法，实现弹出堆顶元素
func (h *intHeap) Pop() any {
	// 待出堆元素存放在最后
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// Len sort.Interface 的方法
func (h *intHeap) Len() int {
	return len(*h)
}

// Less sort.Interface 的方法
func (h *intHeap) Less(i, j int) bool {
	// 如果实现小顶堆，则需要调整为小于号
	return (*h)[i].(int) > (*h)[j].(int)
}

// Swap sort.Interface 的方法
func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Top 获取堆顶元素
func (h *intHeap) Top() any {
	return (*h)[0]
}
