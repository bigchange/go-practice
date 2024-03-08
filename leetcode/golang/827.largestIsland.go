/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/18
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/18 18:36
 */
package leetcode


func largestIsland(grid [][]int) int {
	nums := make([][][]int, len(grid))
	rows := len(grid)
	cols := len(grid[0])

	for i, row := range grid {
		nums[i] = make([][]int, cols)
		for j, _ := range row {
			nums[i][j] = make([]int, 2)
		}
	}
	var getNums func(i, j int) int
	getNums = func(i, j int) int {
		num := 0
		if grid[i][j] != 0 {
			// 向左
			if j > 0 {
				num += nums[i][j - 1][0]
				nums[i][j][0] = nums[i][j - 1][0] + 1
			}
			// 向上
			if i > 0 {
				num += nums[i - 1][j][1]
				nums[i][j][1] = nums[i - 1][j][1] + 1
			}
		}
		return num
	}
	for  i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			getNums(i,j)
		}
	}
  maxNum := 0
	for  i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 如果将这个0变成1
			if grid[i][j] == 0 {
				grid[i][j] = 1
				num := getNums(i, j) + 1
				// 向下
				for k := i+1; k < rows; k++ {
					if grid[k][j] != 1 {
						break
					}
					num++
				}
				// 向右
				for k := j+1; k < cols; k++ {
					if grid[i][k]  != 0 {
						break
					}
					num++
				}
				maxNum = max(maxNum, num)
				// 重现
				grid[i][j] = 0
			}
		}
	}
	return maxNum
}