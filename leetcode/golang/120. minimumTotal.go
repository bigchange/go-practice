/*
 * @Author: Jerry You
 * @CreatedDate: 2021/9/5
 * @Last Modified by: lolaliva
 * @Last Modified time: 2021/9/5 3:48 PM
 */
package leetcode

import "math"

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, i + 1)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		// 当前行的元素个数与当前所在的层数相等
		for j := 0; j < i + 1; j++ {
			dp[i][j] = math.MaxInt64
			val := triangle[i][j]
			if  j == 0 {
				// 第一个元素：只能从上一层相同下标的位置过来
				// 取最小和
				dp[i][j] = min(dp[i][j], dp[i-1][j] + val)
			} else if  j == i {
				// 最后一个元素： 只能从上一层，下标-1的位置过来
				// 取最小和
				dp[i][j] = min(dp[i][j],  dp[i-1][j-1] + val)
			} else {
				// 中间元素: 可以从相同下标，或者下标-1，两个位置过来
				// 取最小和
				dp[i][j] = min(dp[i][j], min(dp[i-1][j] + val, dp[i-1][j-1]+ val))
			}
		}
	}
	// 取最后一层中：数值最小的即为最小和
	ans := math.MaxInt64
	for i := 0; i < m; i++ {
		ans = min(dp[m -1][i], ans)
	}
  return ans
}