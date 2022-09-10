/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/10
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/10 22:33
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
// 深度搜索， 越来越有感觉了哈哈
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 得到左边裁剪后的节点
	leftNode := trimBST(root.Left, low, high)
	// 得到右边边裁剪后的节点
	rightNode := trimBST(root.Right, low, high)

	// 先判断根节点是否符合
	// 1. 根节点符合
	if isEffectNode(root, low, high) {
		// 保留根节点
		root.Left = leftNode
		root.Right = rightNode
	} else {
		// 2. 根节点不符合
		// 3. 判断返回的右节点是否符合
		// 4.右节点符合，返回右节点
		if isEffectNode(rightNode, low, high) {
			root = rightNode
		} else if isEffectNode(leftNode, low, high) {
			// 5. 右节点不符合，判断左阶段是否符合
			// 6.左节点符合，返回左节点
			root = leftNode
		} else {
			// 7. 根，左右都不符合，返回nil
			root = nil
		}
	}
	return root
}

// 判断节点是否满足条件
func isEffectNode(node *TreeNode, low, high int) bool {
	if node == nil {
		return false
	}
	if low <= node.Val && node.Val <= high {
		return  true
	}
	return false
}
