package tree

// 总结
// 广度优先遍历
// 层遍历： 队列
// Levels 二叉树的层序遍历
func Levels(root *TreeNode) []int {
	var vals []int
	if root == nil {
		return vals
	}
	// slice模拟队列（先入先出）
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		size := len(stack)
		for i := 0; i < size; i++ {
			v := stack[i]
			if v.Left != nil {
				stack = append(stack, v.Left)
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

// 前中后序遍历： 栈
// 中序： 一直压入左子树节点，直到null， 弹出栈顶，出栈顺序就是中序遍历结果，然后转向右边节点，如果右节点为空，当前子树完成，否则发现新子树。
// 前序： 根入栈，栈不为空，弹出栈顶，然后先右节点后左节点入栈， 直到栈为空。 出栈顺序就是前序遍历结果。
// 后序： 根入栈， 栈不为空，弹出栈顶， 然后先左后右节点入栈， 直到栈为空。 出栈的逆序就是后序遍历的结果。
// 层序： 根入队列， 队列不为空， 出队列，根据队列的当前长度，移除所有元素，移除的元素的子节点需要再做入队列的操作，如此循环，直到队列为空。

// 特殊性质： 二叉树搜索树的中序遍历是有序的。

// 前中后序遍历： 递归实现（深度优先遍历），先走到尽头，再回溯继续
// 递归过程可分为“递”和“归”两个相反的部分。
// “递”表示开启新方法，程序在此过程中访问下一个节点；
// “归”表示函数返回，代表该节点已经访问完毕
// refer: https://www.hello-algo.com/chapter_tree/binary_tree_traversal/#822
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	// 前序遍历访问节点位置
	dfs(root.Left) // 继续左子树
	// 中序遍历访问节点位置
	// 已访问左子树
	dfs(root.Right) // 继续右子树
	// 后序遍历访问节点位置
	// 左右子树都已访问即将回溯
}
