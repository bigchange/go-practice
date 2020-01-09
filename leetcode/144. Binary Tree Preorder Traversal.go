package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解题思路
// 前序遍历：使用stack先进后出的方式替换递归，依次遍历stack直至stack为空
// 递归：一般都可以用队列或者栈的方式优化
func preorderTraversal(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		item := stack.Front()
		e := item.Value.(*TreeNode)
		stack.Remove(item)
		ret = append(ret, e.Val)
		if e.Right != nil {
			stack.PushFront(e.Right)
		}
		if e.Left != nil {
			stack.PushFront(e.Left)
		}
	}
	return ret
}
