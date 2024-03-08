package leetcode

func checkValidString(s string) bool {
	var lStack []int
	var starStack []int
	for i, r := range  s {
		if r == '(' {
			lStack = append(lStack, i)
		} else if r == '*' {
			starStack = append(starStack, i)
		} else {
			// 遇到了 ')'
			// 优先匹配 '('
			if len(lStack) > 0 {
				lStack = lStack[:len(lStack) - 1]
			} else if len(starStack) > 0  {
				starStack = starStack[:len(starStack) - 1]
			} else {
				return false
			}
		}
	}
	// 字符遍历结束后
	// 再根据两个栈中的数据判断，'*' 用来当做 ')' ，用来匹配剩下的 '('
	// 但是需要满足 '*' 出现的下标 > '(' 出现的下标
	if len(lStack) == 0 {
		return true
	}
	i, j := len(lStack) - 1, len(starStack) - 1
	for ;i >=0 && j >=0; {
		if lStack[i] > starStack[j] {
			return false
		}
		i--
		j--
	}
	// 左括号都匹配完
	return i < 0
}
