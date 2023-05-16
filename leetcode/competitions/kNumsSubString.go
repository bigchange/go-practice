package competitions

import "fmt"

// 满足字符至少出现k次的最长子串长度
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
				newEn := st + index
				left := getLen(st, newEn)
				right := getLen(newEn+1, en)
				max = maxF(maxF(left, right), max)
				matched = false
				break
			}
		}
		if matched {
			maxLen := en - st
			// 特殊处理整个s都满足的情况
			if en == len(sub) && st == 0 {
				max = maxF(max, maxLen)
			}
			return maxLen
		}
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
