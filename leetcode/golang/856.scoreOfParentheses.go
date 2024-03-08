package leetcode

//给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：
//
//
// () 得 1 分。
// AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
// (A) 得 2 * A 分，其中 A 是平衡括号字符串。
//
//
//
//
// 示例 1：
//
// 输入： "()"
//输出： 1
//
//
// 示例 2：
//
// 输入： "(())"
//输出： 2
//
//
// 示例 3：
//
// 输入： "()()"
//输出： 2
//
//
// 示例 4：
//
// 输入： "(()(()))"
//输出： 6
//
//
//
//
// 提示：
//
//
// S 是平衡括号字符串，且只含有 ( 和 ) 。
// 2 <= S.length <= 50
//
//
// Related Topics 栈 字符串 👍 374 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func scoreOfParentheses(s string) int {
	// 往栈里面压入分数
	var stack = []int{0}
	for _, i := range s {
		// 入栈
		if i == '(' {
			// 压入空串的分数0
			stack = append(stack, 0)
		} else {
			// 出栈
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 出栈， 拿到v，计算新的分数
			// 如果A是空字符串（v = 0），（）的分数 = 1
			// 否则 （A）的分数 = 2 * v
			// 将新的分数加到栈顶中去保存
			stack[len(stack)-1] += max(2*v, 1)
		}
	}
	// 最后结果保存在栈顶
	return stack[0]
}
//leetcode submit region end(Prohibit modification and deletion)

