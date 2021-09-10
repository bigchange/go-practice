package leetcode

import (
	"fmt"
)

func Entry() {
	// nums := []int{3,5,10,1,8,2,9,4,11,7,12,6}
	ret := chalkReplacer([]int{3, 4, 1, 2}, 25)

	fmt.Printf("ret:[%v]\n",ret)

	fmt.Printf("leetCode Entry return:%+v",  ret)
}
