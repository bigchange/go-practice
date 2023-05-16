package tree

type aVLTree = TreeNode

// refer: https://www.hello-algo.com/chapter_tree/avl_tree/#_3
// AVL树旋转：旋转操作既能保持树的「二叉搜索树」属性，也能使树重新变为「平衡二叉树」。
// 旋转操作分为四种：
// 1. 右旋、
// 2. 左旋、
// 3. 先右旋后左旋: 即先对 child 执行「右旋」，然后对 node 执行「左旋」。
// 4. 先左旋后右旋: 即先对 child 执行「左旋」，再对 node 执行「右旋」。

// 扩展：为什么红黑树比 AVL 树更受欢迎？
// 红黑树的平衡条件相对宽松，因此在红黑树中插入与删除节点所需的旋转操作相对较少，在节点增删操作上的平均效率高于 AVL 树。

/* 右旋操作 */
func (t *aVLTree) rightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right
	// 以 child 为原点，将 node 向右旋转
	child.Right = node
	node.Left = grandChild
	// 更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)
	// 返回旋转后子树的根节点
	return child
}

/* 左旋操作 */
func (t *aVLTree) leftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left
	// 以 child 为原点，将 node 向左旋转
	child.Left = node
	node.Right = grandChild
	// 更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)
	// 返回旋转后子树的根节点
	return child
}

/* 执行旋转操作，使该子树重新恢复平衡 */
// 通过判断失衡节点的平衡因子以及较高一侧子节点的平衡因子的正负号来确定旋转操作具体属于哪种情况
func (t *aVLTree) rotate(node *TreeNode) *TreeNode {
	// 获取节点 node 的平衡因子
	bf := t.balanceFactor(node)
	// 左偏树
	if bf > 1 {
		if t.balanceFactor(node.Left) >= 0 {
			// 右旋
			return t.rightRotate(node)
		} else {
			// 先左旋后右旋
			node.Left = t.leftRotate(node.Left)
			return t.rightRotate(node)
		}
	}
	// 右偏树
	if bf < -1 {
		if t.balanceFactor(node.Right) <= 0 {
			// 左旋
			return t.leftRotate(node)
		} else {
			// 先右旋后左旋
			node.Right = t.rightRotate(node.Right)
			return t.leftRotate(node)
		}
	}
	// 平衡树，无需旋转，直接返回
	return node
}

/* 插入节点 */
// 需要从这个节点开始，自底向上执行旋转操作，使所有失衡节点恢复平衡。
func (t *aVLTree) insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

/* 删除节点 */
// 类似地，在二叉搜索树的删除节点方法的基础上，需要从底至顶地执行旋转操作，使所有失衡节点恢复平衡。
func (t *aVLTree) remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

/* 递归删除节点（辅助方法） */
func (t *aVLTree) removeHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	/* 1. 查找节点，并删除之 */
	if val < node.Val {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.removeHelper(node.Right, val)
	} else {
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				// 子节点数量 = 0 ，直接删除 node 并返回
				return nil
			} else {
				// 子节点数量 = 1 ，直接删除 node
				node = child
			}
		} else {
			// 子节点数量 = 2 ，则将中序遍历的下个节点删除，并用该节点替换当前节点
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			node.Right = t.removeHelper(node.Right, temp.Val)
			node.Val = temp.Val
		}
	}
	// 更新节点高度
	t.updateHeight(node)
	/* 2. 执行旋转操作，使该子树重新恢复平衡 */
	node = t.rotate(node)
	// 返回子树的根节点
	return node
}

/* 递归插入节点（辅助方法） */
func (t *aVLTree) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewTreeNode(val)
	}
	/* 1. 查找插入位置，并插入节点 */
	if val < node.Val {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		// 重复节点不插入，直接返回
		return node
	}
	// 更新节点高度
	t.updateHeight(node)
	/* 2. 执行旋转操作，使该子树重新恢复平衡 */
	node = t.rotate(node)
	// 返回子树的根节点
	return node
}

/* 获取节点高度 */
func (t *aVLTree) height(node *TreeNode) int {
	// 空节点高度为 -1 ，叶节点高度为 0
	if node != nil {
		return node.Height
	}
	return -1
}

/* 更新节点高度 */
func (t *aVLTree) updateHeight(node *TreeNode) {
	lh := t.height(node.Left)
	rh := t.height(node.Right)
	// 节点高度等于最高子树高度 + 1
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

/* 获取平衡因子 */
// 设平衡因子为 f ，则一棵 AVL 树的任意节点的平衡因子皆满足 -1 <= f <= 1
func (t *aVLTree) balanceFactor(node *TreeNode) int {
	// 空节点平衡因子为 0
	if node == nil {
		return 0
	}
	// 节点平衡因子 = 左子树高度 - 右子树高度
	return t.height(node.Left) - t.height(node.Right)
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
