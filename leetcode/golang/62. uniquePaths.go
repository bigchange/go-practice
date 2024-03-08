/*
 * @Author: Jerry You
 * @CreatedDate: 2021/9/4
 * @Last Modified by: lolaliva
 * @Last Modified time: 2021/9/4 10:03 PM
 */
package leetcode

func uniquePaths(m int, n int) int {
	paths := make([][]int, m)
	for i:= 0; i < m; i++ {
		pathN := make([]int, n)
		paths[i] = pathN
	}
	// 初始化边界：只有一种可能第一行（向右），第一列（向下）
	for i:= 0; i < m; i++ {
		// 第一列
		paths[i][0] = 1
	}
	for i:= 0; i < n; i++ {
		// 第一行
		paths[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 每个位置可能性： 等同于通过向右或者向下移动过来可能性总和
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	return paths[m-1][n-1]
}