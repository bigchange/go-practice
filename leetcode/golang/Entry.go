package leetcode

import (
	"flag"
	"fmt"
	"time"
)

var s string
var n int
var p int

func Entry() {
	// nums := []int{3,5,10,1,8,2,9,4,11,7,12,6}
	flag.IntVar(&n, "n", 1, "input n")
	flag.IntVar(&p, "p", 1, "input p")
	flag.StringVar(&s, "s", "", "input string")
	flag.Parse()
	TODO()
}

func TODO() int {
	t1 := time.Now()
	var res interface{}
	// begin
	cases := [][]int{
		// cases,
	}
	for _, c := range cases {
		res = run(c)
		fmt.Println("res", res, time.Since(t1))
	}
	// end
	return 0
}
func run(c interface{}) int {
	return 0
}
func printTrees(root *TreeNode) {
	rets := printTree(root)
	for _, r := range rets {
		fmt.Println(r)
	}
}
