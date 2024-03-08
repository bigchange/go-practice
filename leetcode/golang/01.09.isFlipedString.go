package leetcode

import "strings"

//字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。
//
// 示例1:
//
//  输入：s1 = "waterbottle", s2 = "erbottlewat"
// 输出：True
//
//
// 示例2:
//
//  输入：s1 = "aa", s2 = "aba"
// 输出：False
//
//
//
//
//
// 提示：
//
//
// 字符串长度在[0, 100000]范围内。
//
//
// 说明:
//
//
// 你能只调用一次检查子串的方法吗？
//
//
// Related Topics 字符串 字符串匹配 👍 167 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 方法一： 枚举法 + 模拟
// 按照s1的每个位置旋转一次，然后与s2比较是否相等
func isFlipedString(s1 string, s2 string) bool {
	if  s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	for i := 1; i <= len(s1) - 1; i++ {
		left := s1[:i]
		right := s1[i:]
		fliped := right + left
		if  fliped == s2 {
			return true
		}
	}
	return false
}

// 方法二：搜素子字符串
// 构造字符串： s = s1 + s1, s包含了所有旋转后的可能子串( 这也能想出来，不过仔细想一下，感觉是容易理解，就是比较难想出来）
// 只需判断s2是否在其中即可
func isFlipedString_2(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1 + s1, s2)
}
//leetcode submit region end(Prohibit modification and deletion)
