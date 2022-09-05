package leetcode


// 入口
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var sameTrees []*TreeNode
	// 所有子树
	nodes := []*TreeNode{root}
	var dfs func(r *TreeNode) *TreeNode
	dfs = func(r *TreeNode) *TreeNode  {
		if r == nil {
			return nil
		}
		if r.Left != nil {
			size := len(nodes)
			for i := 0; i < size; i++ {
				if isSameTree(nodes[i], r.Left) {
					sameTrees = append(sameTrees, r.Left)
				} else {
					nodes = append(nodes, r.Left)
				}
			}
			nodes = append(nodes, r.Left)
			dfs(r.Left)
		}
		if r.Right != nil {
			size := len(nodes)
			for i := 0; i < size; i++ {
				if isSameTree(nodes[i], r.Right) {
					sameTrees = append(sameTrees, r.Right)
				} else {
					nodes = append(nodes, r.Right)
				}
			}
			nodes = append(nodes, r.Right)
			dfs(r.Right)
		}
		return r
	}
	dfs(root)
	return sameTrees

}


// 判断是否相同的子树
func isSameTree(node1 *TreeNode, node2 *TreeNode) bool {
	if node2 == nil && node1 == nil {
		return true
	}
	if node2 == nil || node1 == nil {
		return false
	}
	if node2.Val == node1.Val {
		 return isSameTree(node1.Left, node2.Left) &&
		 	isSameTree(node1.Right, node2.Right)
	}
	return false
}


