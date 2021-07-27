package leetcode

import (
	"fmt"
)

func Entry() {
	// nums := []int{3,5,10,1,8,2,9,4,11,7,12,6}
	ret := findMaximizedCapital(3,0 , []int{1,2,3}, []int{0, 1,2})
	fmt.Println("ret:", ret)
	fmt.Printf("leetCode Entry return:%+v",  ret)
}
