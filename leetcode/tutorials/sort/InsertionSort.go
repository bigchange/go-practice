package sort

// InsertionSort 插入排序
// Worst Case- O(n*n)
// Best Case- O(n) – When the array is already sorted
// Space Complexity of insertion sort is O(1)
func InsertionSort(arr []int) {
	len := len(arr)
	// 外循环： 插入的次数
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