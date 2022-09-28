/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/11
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/11 22:11
 */
package leetcode

import (
	"container/heap"
	"math"
	"sort"
)

// 入口
func mincostToHireWorkers(quality, wage []int, k int) float64 {
	n := len(quality)
	h := make([]int, n)
	for i := range h {
		h[i] = i
	}
	sort.Slice(h, func(i, j int) bool {
		a, b := h[i], h[j]
		// 按此比值升序排列 w[i]/q[i] < w[j]/q[j]
		return quality[a]*wage[b] > quality[b]*wage[a]
	})
	totalq := 0
	q := hp{}
	// 符合条件的k-1个
	for i := 0; i < k-1; i++ {
		totalq += quality[h[i]]
		heap.Push(&q, quality[h[i]])
	}
	ans := 1e9
	// 当我们以某一个工人 x（第k个） 作为一个工资组中权重最高时，工资组中其他人员只需要在权重小于工人 x 的集合中
	// 选择工作质量最少的 k-1 名工人来组成工资组即可
	for i := k - 1; i < n; i++ {
		idx := h[i]
		totalq += quality[idx]
		heap.Push(&q, quality[idx])
		ans = math.Min(ans, float64(wage[idx])/float64(quality[idx])*float64(totalq))
		// 剔除其中工作质量最大的那个
		totalq -= heap.Pop(&q).(int)
	}
	return ans
}