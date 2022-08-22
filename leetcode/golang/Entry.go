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
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},

	}
	ret := printTree(root)

	for _, item := range ret {
		for _, i := range item{
			fmt.Print(i + " ")
		}
		fmt.Println("\n")
	}
}
