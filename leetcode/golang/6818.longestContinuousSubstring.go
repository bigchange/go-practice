package leetcode

func longestContinuousSubstring(s string) (ans int) {
	start := 0
	for i := 1; i < len(s); i++ {
		// 后一位比前一位大1，就是连续的
		if s[i] != s[i-1]+1 {
			// 保留最大
			ans = max(ans, i-start)
			// 如果不连续，记录新起点
			start = i
		}
	}
	return max(ans, len(s)-start)
}