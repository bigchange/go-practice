/*
 * @Author: Jerry You
 * @CreatedDate: 2022/11/5
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/11/5 18:26
 */
package leetcode

import "fmt"

//给你一个以字符串形式表述的 布尔表达式（boolean） expression，返回该式的运算结果。
//
// 有效的表达式需遵循以下约定：
//
//
// "t"，运算结果为 True
// "f"，运算结果为 False
// "!(expr)"，运算过程为对内部表达式 expr 进行逻辑 非的运算（NOT）
// "&(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 与的运算（AND）
// "|(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 或的运算（OR）
//
//
//
//
// 示例 1：
//
// 输入：expression = "!(f)"
//输出：true
//
//
// 示例 2：
//
// 输入：expression = "|(f,t)"
//输出：true
//
//
// 示例 3：
//
// 输入：expression = "&(t,f)"
//输出：false
//
//
// 示例 4：
//
// 输入：expression = "|(&(t,f,t),!(t))"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= expression.length <= 20000
// expression[i] 由 {'(', ')', '&', '|', '!', 't', 'f', ','} 中的字符组成。
// expression 是以上述形式给出的有效表达式，表示一个布尔值。
//
//
// Related Topics 栈 递归 字符串 👍 146 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func parseBoolExpr(expression string) bool {
	var values []rune
	for _, r := range []rune(expression) {
		if r == '(' || r == ',' {
			continue
		}
		if r == ')' {
			l := len(values) - 1
			pos := 0
			for i := l; i > 0; i-- {
				v := values[i]
				if v == '&' || v == '|' || v == '!' {
					break
				}
				pos++
			}
			method := values[l-pos]
			var res = parseBool(values[l])
			count := 0
			for i := l - 1; i > l-pos; i-- {
				count++
				res = cal(method, res, parseBool(values[i]))
			}
			if count == 0 {
				if method == '&' {
					res = cal(method, res, true)
				} else {
					res = cal(method, res, false)
				}
			}
			values = values[:l-pos]
			fmt.Println("res, l - pos, l ", res, l-pos, l+1)
			if l-pos == 0 {
				return res
			} else {
				values = append(values, boolParse(res))
			}
			continue
		}
		values = append(values, r)
	}
	return false
}
func parseBool(r rune) bool {
	if r == 't' {
		return true
	}
	return false
}
func boolParse(b bool) rune {
	if b {
		return 't'
	}
	return 'f'
}
func cal(r rune, res bool, v bool) bool {
	if r == '&' {
		return res && v
	}
	if r == '|' {
		return res || v
	}
	if r == '!' {
		return !res
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
