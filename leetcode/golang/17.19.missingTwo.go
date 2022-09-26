package leetcode

import "sort"

//给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？
//
// 以任意顺序返回这两个数字均可。
//
// 示例 1:
//
// 输入: [1]
//输出: [2,3]
//
// 示例 2:
//
// 输入: [2,3]
//输出: [1,4]
//
// 提示：
//
//
// nums.length <= 30000
//
// Related Topics 位运算 数组 哈希表 👍 123 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func missingTwo(nums []int) []int {
	tmp := 0
	var res []int
	c := 0
	sort.Ints(nums)
	for _, i := range nums {
		gap := i - tmp
		if gap == 1 {
			tmp = i
		} else if  gap == 2 {
			res = append(res, i - 1)
			c = c + 1
			tmp = i
		} else if gap == 3 {
			res = append(res, i - 1, i- 2)
			c = c + 2
			break
		}
	}
	// 是否拿到了2个缺失的数
	if  c != 2 {
		for c < 2 {
			tmp = tmp + 1
			res = append(res,  tmp)
			c++
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

