/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/2
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/2 23:19
 */
package leetcode


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 入口
// 使用深度搜索
func longestUnivaluePath(root *TreeNode) int {
	var ret int
	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		// 左子树长度
		ld := dfs(r.Left)
		// 右子树长度
		rd := dfs(r.Right)
		// 左，右子树临时长度
		ldt, rdt := 0,0
		// 如果当前子树不为空，且值相等
		if r.Left != nil && r.Val == r.Left.Val {
			// 根节点与左节点认为是连通的
			// 左子树临时长度 = 左节点返回长度 + 1
			ldt = ld + 1
		}
		// 同左子树处理
		if r.Right != nil && r.Val == r.Right.Val {
			rdt = rd + 1
		}
		// 计算返回结果，当根与左右都相等的时候
		// 长度可累加，为最长
		// 否则： 要么取ldt，或rdt，或者 0
		ret = max(ret, ldt + rdt)
		// 取：其中最大的返回
		return max(ldt, rdt)
	}
	dfs(root)
	return ret
}
