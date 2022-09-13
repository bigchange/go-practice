package leetcode

import "fmt"

func maximumSwap(num int) int {
	ret := 0
	if num == 0 {
		return num
	}
	mul := 1
	nums  := getNums(num)
	fmt.Println("nums ", nums)
	size := len(nums)
	swap(nums, size - 1)
	fmt.Println("nums 1 ", nums)

	if size >= 2 {
		swap(nums, size - 2)
	}
	fmt.Println("nums 2", nums)

	for i := 0; i < len(nums); i++ {
		ret = ret + nums[i] * mul
		mul = mul * 10
	}
	return ret
}

func swap(nums []int, i int) {
	maxV := nums[0]
	maxVIndex := 0
	for k := 1; k < i; k++ {
		if maxV < nums[k]  {
			maxV = nums[k]
			maxVIndex = k
		}
	}
	if maxV > nums[i] {
		nums[maxVIndex], nums[i] = nums[i], nums[maxVIndex]
	}
}
func getNums(num int) []int {
	var nums []int
	for num > 0 {
		v := num % 10
		num = num / 10
		nums = append(nums, v)
	}
	return nums
}
