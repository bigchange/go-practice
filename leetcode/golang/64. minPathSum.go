/*
 * @Author: Jerry You
 * @CreatedDate: 2021/9/5
 * @Last Modified by: lolaliva
 * @Last Modified time: 2021/9/5 12:01 PM
 */
package leetcode

func minPathSum(grid [][]int) int {
	 m := len(grid)
	 n := len(grid[0])
	 for i := 0; i < m; i++ {
	 	for j:= 0; j < n; j++ {
	 		if i > 0 && j > 0 {
	 			// 复用输入的数组，存储已经计算过得最小路径和的数据
	 			// 每次计算一个格子的最小和时： 向上，向左方向格子的数据都是已经计算过的，不会在改变，所以可以用来直接存储最小和当前状态数据
	 			// 省去额外申请空间存储每个位置的最小和的当前状态数据
	 			grid[i][j] = min(grid[i][j] + grid[i-1][j], grid[i][j] + grid[i][j-1])
			} else if i > 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else if j > 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}
		}
	 }
	 return grid[m-1][n-1]
}