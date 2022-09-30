package leetcode

//编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
//
//
//
// 示例 1：
//
// 输入：
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出：
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
//
//
// 示例 2：
//
// 输入：
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出：
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
//
//
// Related Topics 数组 哈希表 矩阵 👍 97 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int)  {
	i, j := 0,0
	rows := len(matrix)
	cols := len(matrix[0])
	rowM := make(map[int]bool)
	colM := make(map[int]bool)
	for i = 0; i < rows; i++ {
		for j = 0; j < cols; j++ {
			// 记录哪些行与列需要处理为0
			if  matrix[i][j] == 0 {
				rowM[i] = true
				colM[j] = true
			}
		}
	}
	// 处理行
	for i = 0; i < rows; i++ {
		_, ok := rowM[i]
		if ok {
			for j = 0; j < cols; j++ {
				matrix[i][j] = 0
			}
		}
	}
	// 处理列
	for i = 0; i < cols; i++ {
		_, ok := colM[i]
		if ok {
			for j = 0; j < rows; j++ {
				matrix[j][i] = 0
			}
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
