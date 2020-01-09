package leetcode

// 1. 双指针的方式
// 慢指针： 指向数组中不重复的最后一个元素，
// 快指针： 一直往后遍历数组
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	len := 0
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num != nums[len] {
			len++
			nums[len] = num
		}
	}
	return len + 1
}

// 2. 计数排序的思想: 计算每个元素应该向前移动的位数，将需要前移的元素前移，数组元素个数减去最后一个元素需要前移的位数即是非重复元素的个数。
// 还可以实现最多保留 k 次重复，实用性高。
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var count int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}
	return len(nums) - count
}
