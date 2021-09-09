package leetcode

import (
	"fmt"
)

func Entry() {
	// nums := []int{3,5,10,1,8,2,9,4,11,7,12,6}
	ret := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	for _, r := range ret {
		fmt.Printf("ret:[%v], len:%v\n", r, len(r))
	}
	fmt.Printf("leetCode Entry return:%+v",  ret)
}
