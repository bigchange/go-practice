/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/3
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/3 17:00
 */
package leetcode

import "sort"

// 入口
// 思路： 每次选数对，选第二个数中最小的， 这样能使得选后续数对有更大的空间范围
// 所以每次选择能保证其是最优的
// 最后的结果也是最长的可能数对集合
// 所以，可先对每个数对中第二个值进行排序
// 然后依次按条件筛选符合规则的数对即可
func findLongestChain(pairs [][]int) int  {
	nums := 1
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	currentMinVal := pairs[0][1]
	for i, p := range pairs {
		if i == 0 {
			continue
		}
		// 符合规则: 下一个数对第一个值要大于前面选的数对的第二个值
		if currentMinVal < p[0] {
			nums++
			currentMinVal = p[1]
		}
	}
	return nums
}
