package leetcode

import (
	"fmt"
	"math"
)

// 动态规划： 从第一行递推可能转移的过程，计算每行各个位置的结果
// 递推转移： 每行的位置的结果，可由上一行的任意位置转移过来，取转移中最小的即可（需要排除正上方来的，相邻行不能取同一列）
// 最后一行的各个位置的最小值，即为最终答案
func minFallingPathSum1289(matrix [][]int) int {
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
			if i == 0 {
				results[i&1][j] = math.MaxInt
				results[i&1][j] = min(results[i&1][j], val)
			} else {
				index := (i - 1) & 1
				results[i&1][j] = findMinVal(results[index], j) + val
			}
			// 此处直接在递推过程中记录最后一行的最小值
			if i == rowLen-1 {
				finalMinVal = min(finalMinVal, results[i&1][j])
			}
		}
		// fmt.Println("====")
		// print(results)
	}
	// 返回结果： 最后一行的最小值
	return finalMinVal
}

// 查找最小值和次小值（该方法可以考虑优化一下， 临时变量缓存）
func findMinVal(res []int, current int) int {
	min := math.MaxInt
	for k := 0; k < len(res); k++ {
		if res[k] < min && k != current {
			min = res[k]
		}
	}
	return min
}

func minFallingPathSum1289_2(matrix [][]int) int {
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
	// i1 代表最小值列下标，i2 代表次小值列下标
	var i1 = -1
	var i2 = -1
	for i := 0; i < rowLen; i++ {
		// 当前转移第 i 行，使用临时变量保存转移过程中的「最小值列下标」&「次小值列下标」
		var ti1 = -1
		var ti2 = -1
		for j := 0; j < cloLen; j++ {
			val := matrix[i][j]
			if i == 0 {
				results[i&1][j] = math.MaxInt
				results[i&1][j] = min(results[i&1][j], val)
				var tmp, tmp2 int
				// 更新 i1 和 i2
				if i1 == -1 {
					tmp = math.MaxInt
				} else {
					tmp = results[i&1][i1]
				}
				if i2 == -1 {
					tmp2 = math.MaxInt
				} else {
					tmp2 = results[i&1][i2]
				}
				if val < tmp {
					i2 = i1
					i1 = j
				} else if val < tmp2 {
					i2 = j
				}
			} else {
				index := (i - 1) & 1
				// 更新动规值
				// 可以选择「最小值」的列选择「最小值」
				if j != i1 {
					results[i&1][j] = results[index][i1] + val
					// 不能选择「最小值」的列选择「次小值」
				} else {
					results[i&1][j] = results[index][i2] + val
				}
				// 更新 i1 和 i2
				var tmp, tmp2 int
				if ti1 == -1 {
					tmp = math.MaxInt
				} else {
					tmp = results[i&1][ti1]
				}
				if ti2 == -1 {
					tmp2 = math.MaxInt
				} else {
					tmp2 = results[i&1][ti2]
				}
				if val < tmp {
					ti2 = ti1
					ti1 = j
				} else if val < tmp2 {
					ti2 = j
				}
			}
		}
		if i != 0 {
			i1 = ti1
			i2 = ti2
		}
		fmt.Println("====")
		print(results)
	}
	var finalMinVal = math.MaxInt
	for i := 0; i < cloLen; i++ {
		finalMinVal = min(finalMinVal, results[(rowLen-1)&1][i])
	}
	return finalMinVal
}
