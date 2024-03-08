/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/4
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/4 14:27
 */
package leetcode


// 入口
func numSpecial(mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])
	ans := 0
	for i :=0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1  && isSpecial(mat, i,j ){
			 ans++
			}
		}
	}
	return ans
}

// 判断是否为特殊位置
// 第i行，第j列元素都为0
func isSpecial(mat [][]int, i, j int) bool {
	rows := len(mat)
	cols := len(mat[0])
	for k := 0; k < rows; k++ {
		if k == i {
			continue
		}
		if mat[k][j] != 0 {
			return false
		}
	}
	for k := 0; k < cols; k++ {
		if  k == j {
			continue
		}
		if mat[i][k] != 0 {
			return false
		}
	}
	return true
}


