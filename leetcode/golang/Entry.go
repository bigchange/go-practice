package leetcode

import (
	"flag"
	"fmt"
	"time"
)

func Entry() {
	// nums := []int{3,5,10,1,8,2,9,4,11,7,12,6}
	var s string
	var n int
	var p int
	flag.IntVar(&n, "n", 1, "input n")
	flag.IntVar(&p, "p", 1, "input p")
	flag.StringVar(&s, "s", "", "input string")
	flag.Parse()
	t1  := time.Now()
	res := TODO()
	fmt.Println("res", res, time.Since(t1))
}

func TODO() int{
	// begin
	// end
	return 0
}

func printTrees(root *TreeNode) {
	rets := printTree(root)
	for _, r := range rets {
		fmt.Println(r)
	}
}
