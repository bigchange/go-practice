package leetcode

import (
	"fmt"

	"github.com/bigchange/go-practice/leetcode/tutorials"
)

func Entry() {
	nums := []int{3,5,10,1,8,2,9,4,11,7,12,6}
	ret := tutorials.RandomSelect(nums, 8)
	fmt.Println("nums:", nums)
	fmt.Printf("leetCode Entry return:%+v",  ret)
}
