package leetcode

//给定一个字符串 s，计算 s 的 不同非空子序列 的个数。因为结果可能很大，所以返回答案需要对 10^9 + 7 取余 。
//
// 字符串的 子序列 是经由原字符串删除一些（也可能不删除）字符但不改变剩余字符相对位置的一个新字符串。
//
//
// 例如，"ace" 是 "abcde" 的一个子序列，但 "aec" 不是。
//
//
//
//
// 示例 1：
//
//
//输入：s = "abc"
//输出：7
//解释：7 个不同的子序列分别是 "a", "b", "c", "ab", "ac", "bc", 以及 "abc"。
//
//
// 示例 2：
//
//
//输入：s = "aba"
//输出：6
//解释：6 个不同的子序列分别是 "a", "b", "ab", "ba", "aa" 以及 "aba"。
//
//
// 示例 3：
//
//
//输入：s = "aaa"
//输出：3
//解释：3 个不同的子序列分别是 "a", "aa" 以及 "aaa"。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2000
// s 仅由小写英文字母组成
//
//
//
//
// Related Topics 字符串 动态规划 👍 229 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func distinctSubseqII(s string) int {
	const mod int = 1e9 + 7
	ans := 0
	// last 记录字符i上一次出现的位置
	// 一共26个字母
	last := make([]int, 26)
	for i := range last {
		last[i] = -1
	}
	n := len(s)
	// f[i] 表示 以 s[i] 为最后一个字符的子序列数量
	f := make([]int, n)
	for i := range f {
		f[i] = 1
	}
	for i, c := range s {
		for _, j := range last {
			// 前面已出现的字符累积到f[i]
			if j != -1 {
				f[i] = (f[i] + f[j]) % mod
			}
		}
		last[c-'a'] = i
	}

	for _, i := range last {
		// 以所以已出现的字母的位最后一个字符的子序列总数
		if i != -1 {
			ans = (ans + f[i]) % mod
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
