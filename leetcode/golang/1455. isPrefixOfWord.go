package leetcode

import (
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	start := 0
	cnt := 1
	for i, v := range sentence {
		// 按空格分割，遇到空格就是一个词的结束
		if  v == ' ' {
			// fmt.Println("item:", sentence[start:i])
			if strings.HasPrefix(sentence[start:i], searchWord) {
				return cnt
			}
			// 下标加1，继续下一个词
			cnt++
			// 记录下一个单词的起始位置
			start = i + 1
		}
	}
	// 最后一个词的结尾不是空格，需要单独判断
	if strings.HasPrefix(sentence[start:], searchWord) {
		return cnt
	}
	// 没有匹配到结果返回-1
	return  -1
}
