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
