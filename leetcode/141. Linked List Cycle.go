package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针法
// 慢指针： 每次往前走一步
// 快指针： 每次往前走两步
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
