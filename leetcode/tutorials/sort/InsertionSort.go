package sort


// 插入排序
// Worst Case- O(n*n)
// Best Case- O(n) – When the array is already sorted
// Space Complexity of insertion sort is O(1)
func insertionSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}