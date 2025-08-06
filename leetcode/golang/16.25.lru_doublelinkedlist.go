package leetcode

import "fmt"

/**
实现一个固定容量的 LRU (Least Recently Used) 缓存。它应该支持以下方法：

Get(key)：如果键存在，返回其值，并将该键移动到缓存的头部（最近使用）。如果键不存在，返回空值。

Put(key, value)：如果键已存在，更新其值，并将其移动到头部。

如果键不存在且缓存未满，添加键值对到头部。

如果键不存在且缓存已满，则删除最久未使用的键值对，再添加新的键值对。
**/

type LinkNode struct {
	// 保存
	// key和value
	// Next和Before指针
	Key    int
	Value  int
	Next   *LinkNode
	Before *LinkNode
}

type LRUCache_2 struct {
	Cap  int
	keyM map[int]*LinkNode
	head *LinkNode
	tail *LinkNode
}

func moveToHead(head *LinkNode, node *LinkNode) {
	if node.Before == head {
		return
	}
	nodeBefore := node.Before
	nodeNext := node.Next
	// 先断开node的链接
	nodeBefore.Next = nodeNext
	nodeNext.Before = nodeBefore
	// 在把Node放入头部
	node.Before = head
	node.Next = head.Next
	head.Next.Before = node
	head.Next = node

}
func ConstructorLRU_2(capacity int) *LRUCache_2 {
	l := &LRUCache_2{
		Cap: capacity, keyM: make(map[int]*LinkNode),
		head: &LinkNode{Value: -2},
		tail: &LinkNode{Value: -2},
	}
	l.head.Next = l.tail
	l.tail.Before = l.head
	return l
}

func (l *LRUCache_2) Get(key int) int {
	v, ok := l.keyM[key]
	if ok {
		// move to head
		moveToHead(l.head, v)
		fmt.Printf("GET key:%v, Value:%v, head Value:%v\n", key, v.Value, l.head.Next.Value)
		return v.Value
	}
	return -1
}

func (l *LRUCache_2) Put(key int, value int) {
	v, ok := l.keyM[key]
	if ok {
		// 如果存在，可能value不同，需要更新最新的value
		v.Value = value
		moveToHead(l.head, v)
		return
	}
	total := len(l.keyM)
	// 需要evict
	if total == l.Cap {
		evictNode := l.tail.Before
		evictNode.Before.Next = l.tail
		l.tail.Before = evictNode.Before
		// 清空需要evict的next和before指针
		evictNode.Next = nil
		evictNode.Before = nil
		// 清除map中的key
		delete(l.keyM, evictNode.Key)
	}
	// 插入到头部
	node := &LinkNode{Value: value, Key: key}
	node.Before = l.head
	node.Next = l.head.Next
	l.head.Next.Before = node
	l.head.Next = node
	l.keyM[key] = node
	return
}
