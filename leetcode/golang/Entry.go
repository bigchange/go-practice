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
	root3 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	printTrees(root3)
	total  := trimBST(root3, 1,2)
	fmt.Println("results")
	printTrees(total)
}

func printTrees(root *TreeNode) {
	rets :=  printTree(root)
	for _, r := range rets {
		fmt.Println(r)
	}
}
