package leetcode

//我们构建了一个包含 n 行( 索引从 1 开始 )的表。首先在第一行我们写上一个 0。接下来的每一行，将前一行中的0替换为01，1替换为10。
//
//
// 例如，对于 n = 3 ，第 1 行是 0 ，第 2 行是 01 ，第3行是 0110 。
//
//
// 给定行数 n 和序数 k，返回第 n 行中第 k 个字符。（ k 从索引 1 开始）
//
// 示例 1:
//
//
//输入: n = 1, k = 1
//输出: 0
//解释: 第一行：0
//
//
// 示例 2:
//
//
//输入: n = 2, k = 1
//输出: 0
//解释:
//第一行: 0
//第二行: 01
//
//
// 示例 3:
//
//
//输入: n = 2, k = 2
//输出: 1
//解释:
//第一行: 0
//第二行: 01
//
//
//
//
// 提示:
//
//
// 1 <= n <= 30
// 1 <= k <= 2n - 1
//
//
// Related Topics 位运算 递归 数学 👍 212 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func kthGrammar(n int, k int) int {
	// 每一行相对于上一行的偏移量
	offset := make([]int, n + 1)
	// 每一行通过相对于上一行的位置
	pos := make([]int, n + 1)
	pos[n] = k
	offset[n] = pos[n] % 2
	for i := n; i >= 1; i-- {
		// 当前行的偏移量
		offset[i] = pos[i] % 2
		add := 0
		// 是否能被整除
		if offset[i] != 0 {
			add = 1
		}
		// 上一层的位置，可通过当前层位置计算出来
		pos[i - 1] = pos[i] / 2 + add
	}
	// 第一行一定是0
	currentV := 0
	// 根据当前v和下一行对应的偏移量，计算结果
	chooseNext := func(v int, offset int) int {
		// 规则：
		// 1. 当前行v = 0，位置为i, 下一层的值为01（0是在位置2 * i - 1， 1是在位置2 * i）
		// 2. 当前行v = 1， 下一层的值为10（1是在位置2 * i - 1， 0是在位置2 * i）
		// 下一层的位置是通过已经计算的下一层偏移量offset来判断
		if v == 0 {
			return v + (1 - offset)
		}
		return  v - (1 - offset)
	}
	// 从第二行开始计算，直到第n行
	for i := 2; i <= n; i++ {
		currentV = chooseNext(currentV, offset[i])
	}
	return currentV
}
//leetcode submit region end(Prohibit modification and deletion)

