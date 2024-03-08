package leetcode

// 入口
// 方法二： 模拟 + 乘法原理
// 定义 s[i], left[i], right[i]
// left[i] 为： s[i]能够作为子数组唯一字符的最左边界，即s[i]左边对一个与s[i]值相等的位置， 若不存在，left[i] = -1
// right[i] 为： s[i]能够作为子数组唯一字符的最右边界，即s[i]右边对一个与s[i]值相等的位置，若不存在，right[i] = n
// 最后累加： 计算每个s[i]字符，能作为唯一字符的子串个数， (i - left[i]) * (right[i] - i)
// 总结： 对每个字符算出它可能在多少个子串中唯一出现，然后累加起来就可以了
func uniqueLetterString_2(s string) int {
	total := 0
	n := len(s)
	left := make([]int, n)
	right := make([]int, n)

	// 用来记录每个字符s[i], 最左边界
	leftMap := make(map[int32]int)
	// 用来记录每个字符s[i], 最右边界
	rightMap := make(map[int32]int)
	// 初始化
	for _, item := range s {
		leftMap[item] = -1
		rightMap[item] = n
	}

	// 计算leftMap和left[i]
	for i, item := range s {
		left[i] = leftMap[item]
		// 每遇到一个字符，记录位置，如果相同的话，保留并记录最后面出现的位置
		leftMap[item] = i
	}

	// 计算rightMap和right[i]
	i := n
	for i = n - 1; i >= 0; i-- {
		item := int32(s[i])
		// 每遇到一个字符，记录位置，如果相同的话，保留并记录最前面出现的位置
		right[i] = rightMap[item]
		rightMap[item] = i
	}

	for k,_ := range s {
		total += (k - left[k]) * (right[k] - k)
	}
	return total
}

// 方法一： 深度搜索（超出时间限制）
func uniqueLetterString(s string) int {
	total := 0
	var dfs func(subStr string, dep int, size int)
	dfs = func(subStr string, dep int, size int) {
		if dep == size {
			return
		}
		newStr := subStr + string(s[dep])
		total += countUniqueChar(newStr)
		dfs(newStr, dep + 1, size)
	}
	size := len(s)
	// 按照不同位置开始递归，计算可能的子串
	for i := 0; i < size; i++ {
		dfs("", i, size)
	}
	return total
}

// 统计字符串中唯一字符个数
func countUniqueChar(s string) int {
	hashMap := make(map[int32]int)
	total := 0
	for _, item := range s {
		hashMap[item]++
	}
	for _, item := range s {
		v, _ := hashMap[item]
		if v == 1 {
			total++
		}
	}
	return total
}

