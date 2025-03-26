package dp

/* 编辑距离：动态规划 
 * 给定两个字符串 s 和 t，返回将 s 转换成 t 所需的最少操作数。
 * 分别需要确定： s 和 t 在不同长度时，如何进行操作使得 s 转换成 t
 * dp[i][j] 表示 s[0:i] 转换为 t[0:j] 的最少操作数
 * 边界条件1： i = 0 时，dp[0][j] = j， 表示将空s字符串转换成 t[0:j] 需要 j 次操作
 * 边界条件2： j = 0 时，dp[i][0] = i， 表示将 s[0:i] 转换成空字符串 t 需要 i 次操作
 * 状态转移：
	* 1. 若 s[i-1] == t[j-1]，则 dp[i][j] = dp[i-1][j-1]
	* 2. 若 s[i-1] != t[j-1]，则 dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
*/
func editDistanceDP(s string, t string) int {
	n := len(s)
	m := len(t)
	MinInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	// 确定边界条件
	// 状态转移：首行首列
	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= m; j++ {
		dp[0][j] = j
	}
	// 状态转移：其余行列
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				// 若两字符相等，则直接跳过此两字符
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 最少编辑步数 = 插入、删除、替换这三种操作的最少编辑步数 + 1
				// 理解：分别考虑三种操作（插入，删除，替换），使得两字符串最后一位相等，每种操作下
				// 剩余的就是重新定义子问题了， 取其最小值记为最少操作步数，加上1对应本次操作的选择
				dp[i][j] = MinInt(MinInt(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[n][m]
}
