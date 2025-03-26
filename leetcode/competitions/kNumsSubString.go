package competitions

import "fmt"

// 满足字符至少出现k次的最长子串长度
// getMaxSubStrLen 计算字符串s中所有字符至少出现k次的最长子串长度
// 参数s是输入的字符串，参数k是字符出现次数的最小阈值
// 返回值是满足条件的最长子串的长度
func getMaxSubStrLen(s string, k int) int {
	var getLen func(s int, e int) int
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	fmt.Println("map:", m)
	max := 0
	getLen = func(st int, en int) int {
		if st >= en {
			return 0
		}
		sub := s[st:en]
		fmt.Println("sub:", sub, len(sub))
		matched := true
		for index, r := range sub {
			if m[r] < k {
				// 注意：在sub中的index需要加上st，才是该index对应字符在整个s中的下标位置
				newEn := st + index
				// 根据不满足K的字符所在位置，将原字符串分成两部分，递归求解
				// 根据返回的左右两边的长度，取最大值保留到max变量中
				left := getLen(st, newEn)
				right := getLen(newEn+1, en)
				max = maxF(maxF(left, right), max)
				// 标记该sub子串中存在不满足条件的字符
				matched = false
				break
			}
		}
		// 如果该标记为true，表示该sub中的全部字符都满足出现至少k次
		// 返回该sub的长度
		if matched {
			maxLen := en - st
			// 特殊处理整个s都满足的情况
			if en == len(sub) && st == 0 {
				max = maxF(max, maxLen)
			}
			return maxLen
		}
		// 否则返回0
		return 0
	}
	getLen(0, len(s))
	return max
}
func maxF(i, j int) int {
	if i > j {
		return i
	}
	return j
}
