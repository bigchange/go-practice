package leetcode

func sumPrefixScores(words []string) []int {
	m := make(map[int64]int)
	// 为什么是31？
	var p int64 = 31
	var key int64 = 0
	for _, w := range words {
		for _, i := range w {
			key = p * key + int64((i - 'a') + 1)
			m[key] += 1
		}
	}
	key = 0
	ret := make([]int, len(words))
	for i, w := range words {
		sum := 0
		for _, j := range w {
			key = p * key  + int64((j - 'a') + 1)
			sum += m[key]
		}
		ret[i] = sum
	}
	return ret
}
