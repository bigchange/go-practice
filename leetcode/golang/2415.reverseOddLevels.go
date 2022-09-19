package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 深度搜索
// 同时遍历左右子树
func reverseOddLevels(root *TreeNode) *TreeNode {
	dfs(root.Left, root.Right, true)
	return root
}

func dfs(node1 *TreeNode, node2 *TreeNode, isTrue bool) {
	if node1 == nil {
		return
	}
	if isTrue {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
	// 每一次递归，左右子树进入的深度都是一样
	// 同时也是一起返回上一层

	// node1的左节点和node2的右节点交换
	dfs(node1.Left, node2.Right, !isTrue)
	// node1的右节点和node2的左节点交换
	dfs(node1.Right, node2.Left, !isTrue)
}



