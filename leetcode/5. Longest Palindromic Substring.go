package leetcode

// 对于一个子串而言，如果它是回文串，并且长度大于 2，那么将它首尾的两个字母去除之后，它仍然是个回文串。
// 例如对于字符串 ababa，如果我们已经知道 bab 是回文串，
// 那么 ababa 一定是回文串，这是因为它的首尾两个字母都是 a
// 在状态转移方程中，我们从长度较短的字符串向长度较长的字符串进行转移
func longestPalindrome(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	dp := make([][]bool, len)
	// dp[i][j] 表示 s[i..j] 是否是回文串
	// 初始化：所有长度为 1 的子串都是回文串
	for i:= 0; i < len; i++ {
		inside := make([]bool, len)
		inside[i] = true
		dp[i] = inside
	}
	// 先枚举子串长度
	for l:= 2; l <= len; l++ {
		// 枚举左边界，左边界的上限设置可以宽松一些
		for i := 0; i < len; i++ {
			// 由 l 和 i 可以确定右边界，即 j - i + 1 = l 得
			j := l + i - 1
			// 如果右边界越界，就可以退出当前循环
			if j >= len {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j - i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i + 1][j - 1]
				}
			}

			// 当前的字符长度
			currentLen := j - i + 1
			if dp[i][j] && currentLen > maxLen {
				maxLen = currentLen
				begin = i
			}
		}
	}
	return s[begin: maxLen + begin]
}