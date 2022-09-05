package leetcode

import (
	"flag"
	"fmt"
)

func Entry() {
	// nums := []int{3,5,10,1,8,2,9,4,11,7,12,6}
	var s string
	flag.StringVar(&s, "s", "", "input string")
	flag.Parse()
	nodes  := findDuplicateSubtrees(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 1,
			},
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 5,
			},
		},
	})
	for _,i := range nodes {
		fmt.Println(printTree(i))
	}
	fmt.Println(isSameTree(&TreeNode{Val: 1}, &TreeNode{Val: 1}))
}
