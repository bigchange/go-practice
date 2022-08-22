package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


// reverse linked list
func Reverse(headNode *ListNode) *ListNode{
	if headNode ==nil {
		return headNode
	}
	if headNode.Next == nil{
		return headNode
	}
	// 使用newNode一层层往上将反转后的链表头节点返回出去
	// 分析就按照stack的进出逻辑来分析就可以了
	// 所谓的递归就是用一个stack来处理的逻辑
	var newNode = Reverse(headNode.Next)
	headNode.Next.Next = headNode
	headNode.Next = nil
	return newNode
}