package sort

// 冒泡排序
// O(n*n)
// Space Complexity of bubble sort is O(1)
func bubbleSort(arr []int) {
	len := len(arr)
	// 次数
	for i := 0; i < len-1; i++ {
		// 每多排一次，减少一个元素的排序
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

/* 冒泡排序 */
func bubbleSort2(nums []int) {
	// 外循环：待排序元素数量为 n-1, n-2, ..., 1
	// i : 标记j能到达的最大位置，刚开始是最后一位，之后每内循环一次之后，减1
	for i := len(nums) - 1; i > 0; i-- {
		// 内循环：冒泡操作
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

/* 冒泡排序（标志优化）*/
func bubbleSortWithFlag(nums []int) {
	// 外循环：待排序元素数量为 n-1, n-2, ..., 1
	for i := len(nums) - 1; i > 0; i-- {
		flag := false // 初始化标志位
		// 内循环：冒泡操作
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true // 记录交换元素
			}
		}
		if flag == false { // 此轮冒泡未交换任何元素，直接跳出
			break
		}
	}
}
