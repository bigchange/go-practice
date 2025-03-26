package sort

// quickAscendingSort 使用快速排序算法对切片 arr 进行升序排序
// arr: 待排序的整数切片
// start: 排序范围的起始索引
// end: 排序范围的结束索引
func quickAscendingSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		// 哨兵节点
		key := arr[(start+end)/2]
		for i <= j {
			//  第一个比哨兵节点大的位置i
			for arr[i] < key {
				i++
			}
			//  第一个比哨兵节点小的位置j
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickAscendingSort(arr, start, j)
		}
		if end > i {
			quickAscendingSort(arr, i, end)
		}
	}
}

/* 快速排序 */
// quickSort2 使用快速排序算法对切片 nums 进行排序
// 参数：
// nums：待排序的整数切片
// left：当前子数组的左边界索引
// right：当前子数组的右边界索引
func quickSort2(nums []int, left, right int) {
	// 子数组长度为 1 时终止递归
	if left >= right {
		return
	}
	// 哨兵划分
	pivot := partition(nums, left, right)
	// 递归左子数组、右子数组
	quickSort2(nums, left, pivot-1)
	quickSort2(nums, pivot+1, right)
}

/* 快速排序（尾递归优化）*/
// quickSort1 对切片 nums 进行快速排序，left 和 right 分别表示要排序的切片的起始和结束索引。
//
// 参数：
// nums []int: 待排序的整数切片
// left int: 排序范围的起始索引
// right int: 排序范围的结束索引
//
// 返回值：
// 无返回值，直接对输入的切片进行排序。
func quickSort1(nums []int, left, right int) {
	// 子数组长度为 1 时终止
	for left < right {
		// 哨兵划分操作
		pivot := partition(nums, left, right)
		// 对两个子数组中较短的那个执行快排
		// 此优化能使得未返回的递归栈数量明显减少，从而优化空间复杂度
		// 递归时间复杂度：直接由递归深度来决定，也就是当前未返回的递归方法的数量
		if pivot-left < right-pivot {
			quickSort1(nums, left, pivot-1) // 递归排序左子数组
			left = pivot + 1                // 剩余待排序区间为 [pivot + 1, right]
		} else {
			quickSort1(nums, pivot+1, right) // 递归排序右子数组
			right = pivot - 1                // 剩余待排序区间为 [left, pivot - 1]
		}
	}
}

/* 哨兵划分 */
// 哨兵的选择直接可能影响整个排序的效率，
// 所有快排中有专门的针对哨兵划分的优化
// partition 函数将切片 nums 中的元素按大小进行分区，
// 使得 nums[left..i-1] 中的所有元素都不大于 nums[left]，
// 而 nums[i+1..right] 中的所有元素都大于 nums[left]
// 函数返回基准数 nums[left] 的索引。
func partition(nums []int, left, right int) int {
	// 以 nums[left] 作为基准数
	// 注意：当我们以最左端元素为基准数时，必须先“从右往左查找”再“从左往右查找”
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j-- // 从右向左找首个小于基准数的元素
		}
		for i < j && nums[i] <= nums[left] {
			i++ // 从左向右找首个大于基准数的元素
		}
		// 元素交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 将基准数交换至两子数组的分界线
	nums[i], nums[left] = nums[left], nums[i]
	return i // 返回基准数的索引
}
