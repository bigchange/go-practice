/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/17
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/17 22:31
 */
package leetcode

func maxLengthBetweenEqualCharacters(s string) int {
	maxLen := -1
	// 存储相同字符的最小下标
	count := make(map[int32]int)
	for i, item := range s {
		pre, ok := count[item]
		if !ok {
			count[item] = i
		} else {
			maxLen = max(i-pre - 1, maxLen)
		}
	}
	return maxLen
}
