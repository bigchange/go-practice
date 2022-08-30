package leetcode

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		newRoot := &TreeNode{Val: val}
		newRoot.Left = root
		return newRoot
	} else {
		searchNode := root
		retNode := insertIntoMaxTree(searchNode.Right, val)
		root.Right = retNode
	}
	return root
}