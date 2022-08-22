package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var p *ListNode
	p1 := l1
	p2 := l2
	left := 0
	getNode := func(v1, v2 int) (*ListNode) {
		sum := v1 + v2 + left
		val := sum % 10
		left = sum / 10
		return &ListNode{Val: val}
	}
	for p1 != nil || p2 != nil {
		v1, v2 := 0, 0
		if p1 != nil {
			v1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			v2 = p2.Val
			p2 = p2.Next
		}
		newNode := getNode(v1,v2)
		if head == nil {
			head = &ListNode{}
			head.Next = newNode
			p = head.Next
		} else {
			p.Next = newNode
			p = p.Next
		}
	}
	if left > 0 {
		p.Next = getNode(0, 0)
	}
	return head.Next

}