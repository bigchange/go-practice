package tree

import "fmt"

// 二叉搜索树：
// 「二叉搜索树 Binary Search Tree」满足以下条件：
// 1. 对于根节点，左子树中所有节点的值 < 根节点的值  < 右子树中所有节点的值；
// 2. 任意节点的左、右子树也是二叉搜索树，即同样满足条件 1. ；

// 性质：中序遍历序列是升序的
type bstnode struct {
	value int
	left  *bstnode
	right *bstnode
}

type bst struct {
	root *bstnode
}

func (b *bst) Reset() {
	b.root = nil
}

func (b *bst) Insert(value int) {
	b.insertRecursion(b.root, value)
}

// 搜索目标节点
// 通常定义为: search
func (b *bst) Find(value int) error {
	// 方法一： 递归
	node := b.findRecursion(b.root, value)
	// 方法二： 迭代
	// node = b.findIterator(value)
	if node == nil {
		return fmt.Errorf("value: %d not found in tree", value)
	}
	return nil
}

/* 删除节点 */
// 分三种情况：
// 1. 当待删除节点的子节点数量=0时，表示待删除节点是叶节点，可以直接删除。
// 2. 当待删除节点的子节点数量=1时，将待删除节点替换为其子节点即可。
// 3. 当待删除节点的子节点数量=2时，删除操作分为三步：
// 3.1 找到待删除节点在“中序遍历序列”中的下一个节点，记为 tmp ；
// 3.2 在树中递归删除节点 tmp ；
// 3.3 用 tmp 的值覆盖待删除节点的值；
func (bst *bst) remove(num int) {
	cur := bst.root
	// 若树为空，直接提前返回
	if cur == nil {
		return
	}
	// 待删除节点之前的节点位置
	var pre *bstnode = nil
	// 循环查找，越过叶节点后跳出
	for cur != nil {
		if cur.value == num {
			break
		}
		// 记录父节点
		pre = cur
		if cur.value < num {
			// 待删除节点在右子树中
			cur = cur.right
		} else {
			// 待删除节点在左子树中
			cur = cur.left
		}
	}
	// 若无待删除节点，则直接返回
	if cur == nil {
		return
	}
	// 子节点数为 0 或 1
	if cur.left == nil || cur.right == nil {
		var child *bstnode = nil
		// 取出待删除节点的子节点
		if cur.left != nil {
			child = cur.left
		} else {
			child = cur.right
		}
		// 将子节点替换为待删除节点
		if pre.left == cur {
			pre.left = child
		} else {
			pre.right = child
		}
		// 子节点数为 2
	} else {
		// 获取中序遍历中待删除节点 cur 的下一个节点
		tmp := cur.right
		for tmp.left != nil {
			tmp = tmp.left
		}
		// 递归删除节点 tmp
		bst.remove(tmp.value)
		// 用 tmp 覆盖 cur
		cur.value = tmp.value
	}
}

// 方法一： 递归插入
// insertRecursion 向二叉搜索树中递归插入一个节点
//
// 参数:
//   node: 指向当前节点的指针，初始调用时应传入nil
//   value: 要插入的值
//
// 返回值:
//   指向插入节点的指针
func (b *bst) insertRecursion(node *bstnode, value int) *bstnode {
	if b.root == nil {
		// 第一个元素插入，初始化root
		b.root = &bstnode{
			value: value,
		}
		return b.root
	}
	// 返回直接插入的节点
	if node == nil {
		return &bstnode{value: value}
	}
	if value <= node.value {
		node.left = b.insertRecursion(node.left, value)
	}
	if value > node.value {
		node.right = b.insertRecursion(node.right, value)
	}
	return node
}

// 方法二： 迭代插入
// insertIterative 向二叉搜索树中插入一个新的节点
//
// 参数:
//   - node: 指向当前遍历到的节点的指针，初始调用时应传入 nil
//   - value: 要插入的值
func (b *bst) insertIterative(node *bstnode, value int) {
	if b.root == nil {
		b.root = &bstnode{
			value: value,
		}
	}
	if node == nil {
		return
	}
	// Find the terminalNode where to insert the new node
	var terminalNode *bstnode
	for node != nil {
		terminalNode = node
		if value <= node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	if value <= terminalNode.value {
		terminalNode.left = &bstnode{value: value}
	} else {
		terminalNode.right = &bstnode{value: value}
	}
	return
}

// findIterator 在二叉搜索树中查找并返回包含指定值的节点指针
// 如果未找到指定值，则返回 nil
//
// 参数:
//     value int: 要查找的值
//
// 返回值:
//     *bstnode: 包含指定值的节点指针，若未找到则返回 nil
func (b *bst) findIterator(value int) *bstnode {
	node := b.root
	// 循环查找，越过叶节点后跳出
	for node != nil {
		if node.value < value {
			// 目标节点在 cur 的右子树中
			node = node.right
		} else if node.value > value {
			// 目标节点在 cur 的左子树中
			node = node.left
		} else {
			// 找到目标节点，跳出循环
			break
		}
	}
	// 返回目标节点
	return node
}

func (b *bst) findRecursion(node *bstnode, value int) *bstnode {
	if node == nil {
		return nil
	}
	if node.value == value {
		return b.root
	}
	if value < node.value {
		return b.findRecursion(node.left, value)
	}
	return b.findRecursion(node.right, value)
}

func (b *bst) Inorder() {
	b.inorderRecursion(b.root)
}

func (b *bst) inorderRecursion(node *bstnode) {
	if node != nil {
		b.inorderRecursion(node.left)
		fmt.Println(node.value)
		b.inorderRecursion(node.right)
	}
}
