package sort

//

// refer: https://www.hello-algo.com/chapter_sorting/counting_sort/#1171
/* 计数排序 */
// 完整实现，可排序对象，并且是稳定排序
// counter: 不是简单的计数，而是前缀和，即索引 i 处的前缀和 prefix[i] 等于数组前 i 个元素个数总和
// 前缀和具有明确的意义，prefix[num] - 1 代表元素 num 在结果数组 res 中最后一次出现的索引

// 局限性：
// 1. 计数排序只适用于非负整数
// 2. 计数排序适用于数据量大但数据范围较小的情况
func countingSort(nums []int) {
	// 1. 统计数组最大元素 m
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// 2. 统计各数字的出现次数
	// counter[num] 代表 num 的出现次数
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	// 3. 求 counter 的前缀和，将“出现次数”转换为“尾索引”
	// 即 counter[num]-1 是 num 在 res 中最后一次出现的索引
	for i := 0; i < m; i++ {
		counter[i+1] += counter[i]
	}
	// 4. 倒序遍历 nums ，将各元素填入结果数组 res
	// 初始化数组 res 用于记录结果
	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		// 将 num 放置到对应索引处
		res[counter[num]-1] = num
		// 令前缀和自减 1 ，得到下次放置 num 的索引
		counter[num]--
	}
	// 使用结果数组 res 覆盖原数组 nums
	copy(nums, res)
}
