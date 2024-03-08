package tree

// 树： 链表的表示
type TreeNode struct {
	Val    int
	Height int // 节点高度，AVL 树需要使用的字段
	Left   *TreeNode
	Right  *TreeNode
	root   *TreeNode
}

// 树： 还可以使用数组的形式表示（堆排序中使用的表示方式）
// 按照层序遍历的顺序，将节点依次存入数组中，需要显式表明某个节点是否是null节点
// 性质：父节点索引为：i， 那么左节点： 2*i+1， 右节点：2*i+2

// AVL 树: 平衡二叉搜索树
// 「节点高度」是指从该节点到最远叶节点的距离，即所经过的“边”的数量。
// 需要特别注意的是，叶节点的高度为 0 ，而空节点的高度为 -1
