package leetcode

import (
	"flag"
	"fmt"
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
	//q := []int{10, 20, 5}
	//w := []int{70, 50, 30}
	//k := 2
	res := flipLights(n,p)
	fmt.Println("res", res)
}

func printTrees(root *TreeNode) {
	rets := printTree(root)
	for _, r := range rets {
		fmt.Println(r)
	}
}
