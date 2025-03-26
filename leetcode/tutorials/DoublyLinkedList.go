package tutorials

import (
	"fmt"
)

type node struct {
	// 数据
	data string
	// 前向指针
	prev *node
	// 后向指针
	next *node
}
type doublyLinkedList struct {
	len  int
	// （optional）头尾指针: 可考虑分别使用一个空节点来占位，这样就不需要判断头尾指针是否为空
	tail *node
	head *node
	// 或者直接使用内置数据结构list.List当做双向链表的实现
}

func InitDoublyList() *doublyLinkedList {
	return &doublyLinkedList{}
}

// 头部添加
func (d *doublyLinkedList) AddFrontNodeDLL(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
}

// 尾部添加
func (d *doublyLinkedList) AddEndNodeDLL(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	}
	d.len++
}

// 前向遍历： 双向链表
func (d *doublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("value = %v (adress:%p), prev = %v, next = %v\n", temp.data, temp, temp.prev, temp.next)
		temp = temp.next
	}
	fmt.Println()
	return nil
}
func (d *doublyLinkedList) Size() int {
	return d.len
}

// 反转双向链表
func (d *doublyLinkedList) ReverseDLL() {
	currentNode := d.head
	var nextInList *node
	d.head, d.tail = d.tail, d.head
	for currentNode != nil {
		nextInList = currentNode.next
		// 方法一： 使用中间变量交换数据
		// currentNode.next = currentNode.prev
		// currentNode.prev = nextInList
		// 方法二：
		// 注意： golang 这种数据的交换是可以做到不用中间变量的
		currentNode.next, currentNode.prev = currentNode.prev, currentNode.next
		currentNode = nextInList
	}
}
