package leetcode

import "sort"

// 堆（优先队列）
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
// Less 最大堆
func (h *hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Reverse reverse linked list
func Reverse(headNode *ListNode) *ListNode {
	if headNode == nil {
		return headNode
	}
	if headNode.Next == nil {
		return headNode
	}
	// 使用newNode一层层往上将反转后的链表头节点返回出去
	// 分析就按照stack的进出逻辑来分析就可以了
	// 所谓的递归就是用一个stack来处理的逻辑
	var newNode = Reverse(headNode.Next)
	headNode.Next.Next = headNode
	headNode.Next = nil
	return newNode
}
