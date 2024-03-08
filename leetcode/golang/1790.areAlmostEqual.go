package leetcode

//给你长度相等的两个字符串 s1 和 s2 。一次 字符串交换 操作的步骤如下：选出某个字符串中的两个下标（不必不同），并交换这两个下标所对应的字符。
//
// 如果对 其中一个字符串 执行 最多一次字符串交换 就可以使两个字符串相等，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
// 输入：s1 = "bank", s2 = "kanb"
//输出：true
//解释：例如，交换 s2 中的第一个和最后一个字符可以得到 "bank"
//
//
// 示例 2：
//
// 输入：s1 = "attack", s2 = "defend"
//输出：false
//解释：一次字符串交换无法使两个字符串相等
//
//
// 示例 3：
//
// 输入：s1 = "kelb", s2 = "kelb"
//输出：true
//解释：两个字符串已经相等，所以不需要进行字符串交换
//
//
// 示例 4：
//
// 输入：s1 = "abcd", s2 = "dcba"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s1.length, s2.length <= 100
// s1.length == s2.length
// s1 和 s2 仅由小写英文字母组成
//
//
// Related Topics 哈希表 字符串 计数 👍 60 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func areAlmostEqual(s1 string, s2 string) bool {
	i, j := -1, -1
	for idx := range s1 {
		// 不等的时候
		if s1[idx] != s2[idx] {
			if i < 0 {
				// 记录第一个不相等的位置
				i = idx
			} else if j < 0 {
				// 记录第二个不相等的位置
				j = idx
			} else {
				// 超过2个位置不相等的，直接返会false
				return false
			}
		}
	}
	// i < 0 :没有不相等的位置，两个字符串 s1 和 s2是同一个，返回true
	// j >= 0: 至少2个位置不同。如果只有一个位置不等(仅仅i > 0)，直接返回false
	// 再通过判断这两个位置交换后是否和s2相等
	return i < 0 || j >= 0 && s1[i] == s2[j] && s1[j] == s2[i]
}
//leetcode submit region end(Prohibit modification and deletion)
