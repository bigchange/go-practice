package leetcode

func maximumSwap(num int) int {
	ret := 0
	if num == 0 {
		return num
	}
	mul := 1
	nums  := getNums(num)
	size := len(nums)
	for i := size - 1; i >= 0; i-- {
		isSwapped := swap(nums, i)
		if isSwapped {
			break
		}
	}
	for i := 0; i < len(nums); i++ {
		ret = ret + nums[i] * mul
		mul = mul * 10
	}
	return ret
}

// i：交换num的位置
// 找到num中最右边最大的一位
func swap(nums []int, i int) bool {
	maxV := nums[0]
	maxVIndex := 0
	// 取num中最右边最大的
	for k := 1; k < i; k++ {
		if maxV < nums[k]  {
			maxV = nums[k]
			maxVIndex = k
		}
	}
	// 交换
	if maxV > nums[i] {
		nums[maxVIndex], nums[i] = nums[i], nums[maxVIndex]
		return true
	}
	return false
}
// 按位获取num的每位元素放入数组中
func getNums(num int) []int {
	var nums []int
	for num > 0 {
		v := num % 10
		num = num / 10
		nums = append(nums, v)
	}
	return nums
}
