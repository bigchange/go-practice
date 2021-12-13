package sort


// 插入排序
// Worst Case- O(n*n)
// Best Case- O(n) – When the array is already sorted
// Space Complexity of insertion sort is O(1)
func insertionSort(arr []int) {
	len := len(arr)
	// 外循环： 插入的次数
	for i := 1; i < len; i++ {
		// 内循环： [0，i)已经有序了, 选择第i个元素插入的位置
		// 假设插入的位置为： k， 那么[k,i)之间的元素需要往后移动一位
		// 位置i，相当于一个临时
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}