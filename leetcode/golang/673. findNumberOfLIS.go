package leetcode

import (
	"strconv"
)

// TODO(cjyou): 超出时间限制
// 思考： 如果优化，进行剪枝？
func findNumberOfLIS(nums []int) int {
	l := len(nums)
	keyMap := make(map[string]bool)
	maxLenMap := make(map[int]int)
	// maxLenIndexMap := make(map[int][]int)
	maxLen := 0
	keyFunc := func(indexs []int) string {
		r := ""
		for _, i := range indexs {
			r += strconv.Itoa(i)
		}
		return  r
	}
	max := func(x int, y int) int {
		if x > y {
		return x
		}
		return y
	}
	var dfs func(int,  []int, []int)
	dfs = func(index int, rets []int, indexs []int) {
		lastItem := rets[len(rets) - 1]
		if index == len(nums) {
			k :=  keyFunc(indexs)
			maxLen = max(maxLen, len(indexs))
			if _, ok := keyMap[k]; !ok {
				keyMap[k] = true
				maxLenMap[len(indexs)]++
			}
			// fmt.Printf("index:%v, rets: %v, maxLen:%v, map:%v\n", indexs, rets, maxLen, maxLenMap)
			return
		}
		if nums[index] > lastItem {
			dfs(index+1, append(rets, nums[index]), append(indexs, index))
		}
		if l - index < (maxLen - len(rets)) {
			return
		}
		dfs(index + 1, rets, indexs)

}
	// eg: [1,1,1,2,2,2,3,3,3]
	// 0,1
	// 2,3,4
	// 5,6
	// eg: [-100,-100,0,0,0,100,100,0,0]
	for i := 0; i <= len(nums) - maxLen; i++ {
		if i > 0 &&  nums[i] == nums[i - 1] {
			continue
		}
		var rets []int
		var indexs []int
		rets = append(rets, nums[i])
		indexs = append(indexs, i)
		dfs(i,  rets, indexs)
	}
	return maxLenMap[maxLen]
}
