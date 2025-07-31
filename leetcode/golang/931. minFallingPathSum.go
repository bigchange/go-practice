package leetcode

import "math"

// 动态规划： 从第一行递推可能转移的过程，计算每行各个位置的结果
// 递推转移： 每行的位置的结果，可由正上方，右上方，左上方转移过来，取转移中最小的即可
// 最后一行的各个位置的最小值，即为最终答案
func minFallingPathSum(matrix [][]int) int {
	row := matrix[0]
	rowLen := len(matrix)
	cloLen := len(row)
	// 动态规划中：需要记录的结果
	// 使用滚动数组的方式，
	// 只需要记录每行中的上一行的数据，即可递推算出当前行的每个位置的结果
	results := make([][]int, 2)
	for i := 0; i < 2; i++ {
		results[i] = make([]int, len(row))
	}
	var finalMinVal = math.MaxInt
	for i := 0; i < rowLen; i++ {
		for j := 0; j < cloLen; j++ {
			val := matrix[i][j]
			// 第一行的结果直接赋值（或者使用copy）
			// 然后i从1开始处理
			// copy(results[0&1], matrix[0])
			if i == 0 {
				results[i&1][j] = math.MaxInt
				results[i&1][j] = min(results[i&1][j], val)
			} else {
				index := (i - 1) & 1
				// 正上方
				var minVal = results[index][j]
				// 左上方(边界判断)
				if j-1 >= 0 {
					minVal = min(minVal, results[index][j-1])
				}
				// 右上方（边界判断）
				if j+1 < cloLen {
					minVal = min(minVal, results[index][j+1])
				}
				// 取正上、左上、右上中的最小值、与当前位置的值累加，即为当前位置的最小结果
				results[i&1][j] = minVal + val
			}
			// 此处直接在递推过程中记录最后一行的最小值
			if i == rowLen-1 {
				finalMinVal = min(finalMinVal, results[i&1][j])
			}
		}
	}
	// 返回结果： 最后一行的最小值
	return finalMinVal
}
