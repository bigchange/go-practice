package leetcode

import "strings"

//给你两个字符串数组 word1 和 word2 。如果两个数组表示的字符串相同，返回 true ；否则，返回 false 。
//
// 数组表示的字符串 是由数组中的所有元素 按顺序 连接形成的字符串。
//
//
//
// 示例 1：
//
//
//输入：word1 = ["ab", "c"], word2 = ["a", "bc"]
//输出：true
//解释：
//word1 表示的字符串为 "ab" + "c" -> "abc"
//word2 表示的字符串为 "a" + "bc" -> "abc"
//两个字符串相同，返回 true
//
// 示例 2：
//
//
//输入：word1 = ["a", "cb"], word2 = ["ab", "c"]
//输出：false
//
//
// 示例 3：
//
//
//输入：word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= word1.length, word2.length <= 10³
// 1 <= word1[i].length, word2[i].length <= 10³
// 1 <= sum(word1[i].length), sum(word2[i].length) <= 10³
// word1[i] 和 word2[i] 由小写字母组成
//
//
// Related Topics 数组 字符串 👍 64 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1,"") == strings.Join(word2,"")
}

// 双指针遍历
func arrayStringsAreEqual_2(word1 []string, word2 []string) bool {
	var i,j,ii,jj int
	for i < len(word1) && j < len(word2) {
		if word1[i][ii] != word2[j][jj] {
			return false
		}
		ii++
		if  ii == len(word1[i]) {
			i++
			ii = 0
		}
		jj++
		if  jj == len(word2[j]) {
			j++
			jj = 0
		}
	}
	return i == len(word1) && j == len(word2)
}
//leetcode submit region end(Prohibit modification and deletion)

