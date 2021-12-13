package tutorials

// 找出数组中的topK，
//
// 算法： 随机选择： O(n) 时间复杂度
func RandomSelect(nums []int, topK int) int {
	l := len(nums)
	ret := findMaxKNums(nums, 0, l-1, topK, partition)
	return ret
}

// {5,3,7,1,8,2,9,4,7,7,2,6}
// 左边元素大于参照元素，右边元素小于参照元素，返回参照元素的下标位置
func partition(nums []int, s, e int) int {
	partitionItem := nums[s]
	index := s
	nums[e], nums[s] = nums[s], nums[e]
	for i := s ; i < e; i++ {
		if nums[i] > partitionItem {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	nums[index], nums[e] = nums[e], nums[index]
	return index
}

// 找数组中第k大的元素，两个相同数值的元素当作两个数处理
// 返回第k大的元素值
func findMaxKNums(nums []int, s int, e int, k int, partition func([]int, int, int) int) int  {
	if k > (e - s) {
		panic("k should be less than array size ")
	}
	if k <= 0 {
		return nums[s - 1]
	}
	if s >= e {
		return nums[s]
	}
	// 以数据第一个元素为参照元素
	// 返回下标： 该下标左边的元素大于参照元素，右边包括小于参照元素，
	i := partition(nums, s, e)
	// 参照元素左边的元素个数
	getNums := i - s
	// 刚好有k个，那么直接返回最后一个元素的下标或者元素值本身
	if getNums == k {
		return nums[i - 1]
	}
	// 个数大于k
	if getNums > k {
		// 那么k个元素依旧在最左边
		return findMaxKNums(nums, s, i - 1, k, partition)
	}
	// 少于k， 那么在右边再去找少了看k - getNums - 1个元素出来即可
	return findMaxKNums(nums, i + 1, e, k - getNums - 1, partition)
}
