package sort

// 选择排序
// O(n*n)，
// Space Complexity of selection sort is O(1)
func selectionSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		// minIndex 为当前最小的位置标记
		// 从minIndex之后的元素中挑出最小的，放入minIndex
		// 如此循环 len - 2 次整个数组就有序了
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
}