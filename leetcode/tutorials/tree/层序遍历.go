package tree

// Levels 二叉树的层序遍历
func Levels(root *TreeNode) []int {
	var vals []int
	if root == nil {
		return vals
	}
	// 栈
	stack :=[]*TreeNode{root}
	for len(stack) > 0 {
		size := len(stack)
		for i:= 0; i < size; i++ {
			v := stack[i]
			if v.Left != nil {
				stack = append(stack,v.Left)
			}
			if v.Right != nil {
				stack = append(stack, v.Right)
			}
			vals = append(vals, stack[i].Val)
		}
		stack = stack[size:]
	}
	return vals
}