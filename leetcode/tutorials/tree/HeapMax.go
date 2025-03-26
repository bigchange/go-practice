package tree

import "fmt"

// 自实现Heap

// refer: https://www.hello-algo.com/chapter_heap/build_heap/
/* 构造方法，根据切片建堆，已有元素存储在切片中
 * 时间复杂度： O(n)
 */
func NewMaxHeap(nums []any) *maxHeap {
	// 将列表元素原封不动添加进堆
	h := &maxHeap{data: nums}
	for i := len(h.data) - 1; i >= 0; i-- {
		// 堆化除叶节点以外的其他所有节点
		h.siftDown(i)
	}
	return h
}
// 数据是一个一个依次加入到heap中的
type maxHeap struct {
	data []any
}

/* 获取左子节点索引 */
func (h *maxHeap) left(i int) int {
	return 2*i + 1
}

/* 获取右子节点索引 */
func (h *maxHeap) right(i int) int {
	return 2*i + 2
}

/* 获取父节点索引 */
func (h *maxHeap) parent(i int) int {
	// 向下整除
	return (i - 1) / 2
}

/* 访问堆顶元素 */
func (h *maxHeap) peek() any {
	return h.data[0]
}



// 元素入堆： 一个一个依次入堆 , O(nlogn)
// 给定元素 val ，我们首先将其添加到堆底。添加之后，由于 val 可能大于堆中其他元素，堆的成立条件可能已被破坏。
// 因此，需要修复从插入节点到根节点的路径上的各个节点，这个操作被称为「堆化 Heapify」。
// 考虑从入堆节点开始，从底至顶执行堆化
func (h *maxHeap) push(val any) {
	// 添加节点
	h.data = append(h.data, val)
	// 从底至顶堆化
	h.siftUp(len(h.data) - 1)
}

/* 元素出堆 */
// 1. 交换堆顶元素与堆底元素（即交换根节点与最右叶节点）
// 2. 交换完成后，将堆底从列表中删除
// 3. 从根节点开始，从顶至底执行堆化；从顶至底堆化的操作方向与从底至顶堆化相反
func (h *maxHeap) pop() any {
	// 判空处理
	if h.isEmpty() {
		fmt.Println("error")
		return nil
	}
	// 交换根节点与最右叶节点（即交换首元素与尾元素）
	h.swap(0, h.size()-1)
	// 删除节点
	val := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	// 从顶至底堆化
	h.siftDown(0)

	// 返回堆顶元素
	return val
}

/* 从节点 i 开始，从顶至底堆化 */
func (h *maxHeap) siftDown(i int) {
	for {
		// 判断节点 i, l, r 中值最大的节点，记为 max
		l, r, max := h.left(i), h.right(i), i
		if l < h.size() && h.data[l].(int) > h.data[max].(int) {
			max = l
		}
		if r < h.size() && h.data[r].(int) > h.data[max].(int) {
			max = r
		}
		// 若节点 i 最大或索引 l, r 越界，则无需继续堆化，跳出
		if max == i {
			break
		}
		// 交换两节点
		h.swap(i, max)
		// 循环向下堆化
		i = max
	}
}

/* 从节点 i 开始，从底至顶堆化 */
func (h *maxHeap) siftUp(i int) {
	for {
		// 获取节点 i 的父节点
		p := h.parent(i)
		// 当“越过根节点”或“节点无需修复”时，结束堆化
		if p < 0 || h.Less(i, p) {
			break
		}
		// 交换两节点
		h.swap(i, p)
		// 循环向上堆化
		i = p
	}
}

func (h *maxHeap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *maxHeap) size() int {
	return len(h.data)
}

func (h *maxHeap) Less(i, j int) bool {
	val := h.data[i]
	switch val.(type) {
	case float64:
		return h.data[i].(float64) <= h.data[j].(float64)
	case float32:
		return h.data[i].(float32) <= h.data[j].(float32)
	case int32:
		return h.data[i].(int32) <= h.data[j].(int32)
	case int64:
		return h.data[i].(int64) <= h.data[j].(int64)
	case int16:
		return h.data[i].(int16) <= h.data[j].(int16)
	case int8:
		return h.data[i].(int8) <= h.data[j].(int8)
	case int:
		return h.data[i].(int) <= h.data[j].(int)
	default:
		panic("not support data type")
	}
}

// 交换节点数据
func (h *maxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
