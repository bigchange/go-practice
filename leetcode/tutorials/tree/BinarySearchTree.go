package tree

import "fmt"

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

// 方法一： 递归插入
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

func (b *bst) Find(value int) error {
	node := b.findRecursion(b.root, value)
	if node == nil {
		return fmt.Errorf("value: %d not found in tree", value)
	}
	return nil
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