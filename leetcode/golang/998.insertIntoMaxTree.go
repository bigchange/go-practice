package leetcode

// 入口
// 递归
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	// 如果遇到nil节点，返回val对应的新节点
	if root == nil {
		return &TreeNode{Val: val}
	}
	// val比遇到的节点值大
	if root.Val < val {
		// 新建一个节点
		newRoot := &TreeNode{Val: val}
		// 新节点的左节点为root
		newRoot.Left = root
		return newRoot
	} else {
		// val比遇到的节点值小
		searchNode := root
		// 从该节点的右孩子往下递归查询
		// 返回对应的新子树根节点
		retNode := insertIntoMaxTree(searchNode.Right, val)
		// 将其右子树设置为返回的新子树根结点
		root.Right = retNode
	}
	return root
}