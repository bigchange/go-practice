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
	ret := insertIntoMaxTree(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}, 6)
	for _, i := range printTree(ret) {
		fmt.Println(i)
	}

}
