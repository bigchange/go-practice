/*
 * @Author: Jerry You
 * @CreatedDate: 2021/9/5
 * @Last Modified by: lolaliva
 * @Last Modified time: 2021/9/5 11:43 AM
 */
package leetcode


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	paths := make([][]int, m)
	for i:= 0; i < m; i++ {
		pathN := make([]int, n)
		paths[i] = pathN
	}
	// 初始化边界：只有一种可能第一行（向右），第一列（向下）
	for i:= 0; i < m; i++ {
		// 第一列： 遇到障碍物，那么后面的格子都走不到了
		if obstacleGrid[i][0] == 1 {
			break
		}
		paths[i][0] = 1
	}
	for i:= 0; i < n; i++ {
		// 第一行： 同理第一列
		if obstacleGrid[0][i] == 1 {
			break
		}
		paths[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 遇到障碍物： 跳过
			if obstacleGrid[i][j] == 1 {
				continue
			}
			// 每个位置可能性： 等同于通过向右或者向下移动过来可能性总和
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	return paths[m-1][n-1]
}
