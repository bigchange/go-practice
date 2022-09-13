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
	//q := []int{10, 20, 5}
	//w := []int{70, 50, 30}
	//k := 2
	res := maximumSwap(98368)
	fmt.Println("res", res)
}

func printTrees(root *TreeNode) {
	rets := printTree(root)
	for _, r := range rets {
		fmt.Println(r)
	}
}
