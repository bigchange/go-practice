/*
 * @Author: Jerry You
 * @CreatedDate: 2022/8/28
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/8/28 18:22
 */
package leetcode

func widthOfBinaryTree(root *TreeNode) int {
  ret := 1
  // 给节点编号，假设父节点编号index = i，
  // 其左节点的编号为 2 * i，其右节点的编号为： 2 * i + 1
  ques := []*TreeNode{root}
  nodePos := []int{1}
	for ques != nil {
		size := len(nodePos)
		// 每一层最大宽度： 最右边节点的index - 最左边的节点index + 1
		ret = max(ret, nodePos[size - 1] - nodePos[0] +1 )
		// 层序遍历
		// 每一层节点出列，同时将其子节点入列
		tmpQue := ques
		tmpPos := nodePos
		// 重置每一层队列和对应的节点编号数组
		ques = nil
		nodePos = nil
		for i, node := range tmpQue {
			// 左节点入队列
			if node.Left != nil {
				ques = append(ques, node.Left)
				nodePos = append(nodePos, tmpPos[i] * 2)
			}
			// 右节点入队列
			if  node.Right != nil {
				ques = append(ques, node.Right)
				nodePos = append(nodePos, tmpPos[i] * 2 + 1)
			}
		}
	}
	return ret
}