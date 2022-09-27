package leetcode

import "sort"

//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
//
// 示例 1：
//
// 输入: s1 = "abc", s2 = "bca"
//输出: true
//
//
// 示例 2：
//
// 输入: s1 = "abc", s2 = "bad"
//输出: false
//
//
// 说明：
//
//
// 0 <= len(s1) <= 100
// 0 <= len(s2) <= 100
//
//
// Related Topics 哈希表 字符串 排序 👍 103 👎 0

// CheckPermutation
// leetcode submit region begin(Prohibit modification and deletion)
// 方法一： 哈希表
// 直接统计s1和s2中每个字符的个数，然后比较每个字符的出现次数是否相等
// 为了减少内存的使用，可以只用一个map来比较
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return  false
	}
	m1 := make(map[uint8]int)
	i := 0
	for i < len(s1) {
		m1[s1[i] - 'a']++
		i++
	}
	j := 0
	for j < len(s2) {
		m1[s2[j] - 'a']--
		if m1[s2[j] - 'a'] < 0 {
			return false
		}
		j++
	}
	return true
}

// CheckPermutation_2
// 方法二： 排序
// 将s1和s2排序后，然后依次比较每一位是否相同
func CheckPermutation_2( s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return  false
	}
	var s11 = []byte(s1)
	var s22 = []byte(s2)
	sort.Slice(s11, func(i, j int) bool {
		return s11[i] < s11[j]
	})
	sort.Slice(s22, func(i, j int) bool {
		return s22[i] < s22[j]
	})
	i := 0
	for i < len(s22) {
		if s22[i] != s11[i] {
			return false
		}
		i++
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)