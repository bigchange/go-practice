package leetcode

import (
	"flag"
	"fmt"
)

func Entry() {
	// nums := []int{3,5,10,1,8,2,9,4,11,7,12,6}
	var s string
	flag.StringVar(&s, "s","","input string")
	flag.Parse()
	ret := numberOfBoomerangs([][]int{{0,0}, {1,0}, {-1, 0}, {0,1}, {0, -1}})

	fmt.Printf("ret:[%v]\n",ret)

	fmt.Printf("leetCode Entry return:%v \n",  ret)
}
