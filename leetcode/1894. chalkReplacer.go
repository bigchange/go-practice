package leetcode

import "fmt"

func chalkReplacer(chalk []int, k int) int {

	// 先计算一下所有学生消耗粉笔数的总和
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	// k < sum
	// 这一轮就会消耗完所有粉笔
	getIndex := func(leftK int) int {
		current := 0
		for i, v := range chalk {
			current += v
			if current > leftK {
				return i
			}
		}
		return 0
	}
	if k < sum {
		return getIndex(k)
	}
	// 如果k > sum
	// 计算剩余的数量
	leftK := k % sum
	fmt.Printf("sum:%v, left:%v \n", sum, leftK)
	return getIndex(leftK)
}