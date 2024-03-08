package leetcode

import (
	"math"
)

// 解题思路
// 双指针法: 类似快慢指针的思想，快慢指针也常用于解决链表判环、相遇问题等
// 通过左右两边指针的移动，来优化暴力解法中只固定一边的做法
// O(n)
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max := 0.0
	for l < r {
		max = math.Max(max, math.Min(float64(height[l]), float64(height[r]))*float64((r-l)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return int(max)
}
