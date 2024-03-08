package leetcode

// 深度搜索
func maxProduct(s string) int {
	ans := 0
	// 检查字符串是否是回文串
	check := func(s []rune) bool {
		l := len(s)
		l, r := 0, l - 1
		for l < r {
			if s[l] == s[r] {
				l++
				r--
			} else {
				return false
			}
		}
		return true
	}
	// 深度搜索
	var dfs func(s string, ret1, ret2 []rune, index int)
	dfs = func(s string, ret1, ret2 []rune, index int) {
		if check(ret1) && check(ret2) {
			ans = max(ans, len(ret1) * len(ret2))
		}
		if index == len(s) {
			return
		}
		// ret1使用该字符串
		dfs(s, append(ret1, rune(s[index])), ret2, index + 1)
		// ret2使用该字符串
		dfs(s, ret1, append(ret2, rune(s[index])), index + 1)
		// 都不使用该字符串
		dfs(s, ret1, ret2, index+1)
	}
	// 开始深度搜索
	dfs(s, []rune{}, []rune{}, 0)
	return ans
}