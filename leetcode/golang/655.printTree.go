package leetcode

import (
	"math"
	"strconv"
)

// 入口
func printTree(root *TreeNode) [][]string {
	h := getTreeHeight(root)
	height := h - 1
	m := height + 1
	n := int(math.Pow(2, float64(height+1))) - 1
	// 初始化矩阵
	res := make([][]string, m)
	for i := 0; i < m; i++ {
		res[i] = make([]string, n)
		for j:=0; j < n; j++ {
			res[i][j] = "."
		}
	}
	// 填充矩阵
	rootR := 0
	rootC := (n - 1) / 2
	// 递归填充矩阵， 从根结点所在的行开始
	fillTreeNode(root, res, rootR, rootC, height)
	return res
}

// 填充矩阵，从根结点所在的行开始
// 参数： 当前节点所在行，列坐标，以及高度
// 当前节点左孩子对应坐标计算： getPos(row, col, height, false)
// 当前节点右孩子对应坐标计算： getPos(row, col, height, true)
// 依次从根节点递归出发填充对应坐标位置为节点的val即可
// note： 也可以通过层序遍历方式，填充节点，只是填充的顺序有所差别
func fillTreeNode(root *TreeNode, res [][]string, row, col, height int) {
	// 递归退出条件
	if root == nil {
		return
	}
	// 当前节点
	res[row][col] = strconv.FormatInt(int64(root.Val), 10)
	// 左孩子
	newR, newC := getPos(row, col, height, false)
	fillTreeNode(root.Left, res, newR, newC, height)
	// 右孩子
	newR, newC = getPos(row, col, height, true)
	fillTreeNode(root.Right, res, newR, newC, height)
}

func getPos(row, col, height int, right bool) (int, int) {
	newRow := row + 1
	offset := int(math.Pow(2, float64(height-row-1)))
	newCol := col - offset
	if right {
		newCol = col + offset
	}
	return newRow, newCol
}

// 获取树的高度
func getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getTreeHeight(root.Left), getTreeHeight(root.Right))
}
