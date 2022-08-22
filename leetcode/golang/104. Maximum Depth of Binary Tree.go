package leetcode

import "container/list"

// 1. recursive solution
// 比较容易想出来的方式
// 直观上: 就是递归的去看当前节点上左节点和右节点的深度，取最大，然后 + 1 就是整个树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 2. iterative solution
// 有点按层遍历的感觉， 就是一层一层输出，
// 每输出一层的时候，把该层的子节点入队列
// 每层的结果要全部出队列，一层结束，深度+1， 直到队列为空
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := 0
	queque := list.New()
	queque.PushBack(root)
	for queque.Len() != 0 {
		currentLevelNodes := queque.Len()
		for i := 0; i < currentLevelNodes; i++ {
			item := queque.Front()
			queque.Remove(item)
			node := item.Value.(*TreeNode)
			if node.Left != nil {
				queque.PushBack(node.Left)
			}
			if node.Right != nil {
				queque.PushBack(node.Right)
			}
		}
		l++
	}
	return l
}
