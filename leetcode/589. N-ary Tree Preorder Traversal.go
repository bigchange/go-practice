package leetcode

import (
	"container/list"
)

type Node struct {
	Val      int
	Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 解题思路
// 前序遍历：使用stack先进后出的方式替换递归，依次遍历stack直至stack为空
// 递归：一般都可以用队列或者栈的方式优化
func preorder(root *Node) []int {
	l := list.New()
	var ret []int
	if root == nil {
		return ret
	}
	l.PushBack(root)
	for l.Len() != 0 {
		e := l.Front()
		item := e.Value.(*Node)
		ret = append(ret, item.Val)
		l.Remove(e)
		len := len(item.Children)
		for i := len - 1; i >= 0; i-- {
			l.PushFront(item.Children[i])
		}
	}
	return ret
}
