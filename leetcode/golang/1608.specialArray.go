/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/12
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/12 19:19
 */
package leetcode

import "sort"

func specialArray(nums []int) int {
	// 找出所有的整数
	filterNums := []int{}
	for  _, i := range nums {
		if  i >= 0 {
			filterNums = append(filterNums, i)
		}
	}
	size := len(filterNums)
	if size == 0 {
		return 0
	}
	// 升序排列得到的数据
	sort.Slice(filterNums, func(i, j int) bool {
		return filterNums[i] < filterNums[j]
	})
	// 得到备选的最大x为filterNums的个数
	candidates := make([]int, size)
	j := 0
	for i := size; i > 0; i-- {
		candidates[j] = i
		j++
	}
	// 从最大的x检查，返回是否符合规则
	// 该x必须小于对应位置的filterNum[i] 且 小于filterNum[i - 1]
	// 这样才能满足 >= x 的个数刚好等于 x
	// 这样的x只有一个，符合就直接返回即可
	for i, c := range  candidates {
		if  i == 0 {
			if c <= filterNums[i] {
				return c
			}
		} else {
			if filterNums[i - 1] < c && c <= filterNums[i] {
				return c
			}
		}
	}
	// 没有找到对应的x，返回 -1
	return  -1
}