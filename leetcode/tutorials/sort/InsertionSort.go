package sort

// InsertionSort 插入排序
// Worst Case- O(n*n)
// Best Case- O(n) – When the array is already sorted
// Space Complexity of insertion sort is O(1)
func InsertionSort(arr []int) {
	len := len(arr)
	// 外循环： 选择插入的元素下标，从i=1开始
	for i := 1; i < len; i++ {
		// 内循环： [0，i)已经有序了,
		// 从[0, i-1]中选择第i个元素插入的位置
		for j := 0; j < i; j++ {
			// 第i个元素的插入位置假设为k
			// 0 <= k <= i， k与i位置的元素先交换
			// 然后在有序区间[K+1, i-1]中的元素Y都需要后移一位
			// 通过依次遍历有序区间，将其与i位置元素依次交换，最后得到的效果就是区间元素整体后移了一位
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

/* 插入排序 */
func InsertionSort2(nums []int) {
	// 外循环：待排序元素数量为 n-1, n-2, ..., 1
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		// 内循环：将 base 插入到左边的正确位置
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j] // 1. 将 nums[j] 向右移动一位
			j--
		}
		nums[j+1] = base // 2. 将 base 赋值到正确位置
	}
}
