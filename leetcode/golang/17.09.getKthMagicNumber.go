package leetcode

import (
	"container/heap"
	"sort"
)

//有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。
// 注意，不是必须有这些素因子，而是必须不包含其他的素因子。
// 例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
//
// 示例 1:
//
// 输入: k = 5
//
//输出: 9
//
//
// Related Topics 哈希表 数学 动态规划 堆（优先队列） 👍 146 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 方法一： 优先队列（最小堆）
// 依次从堆顶取出第K个数
// 每次取出的数，分别与3，5，7三个因子相乘，将得到的数加入堆中即可
func getKthMagicNumber(k int) int {
	factors := []int{3, 5, 7}
	h := &hp{sort.IntSlice{1}}
	m := make(map[int]bool)
	for i := 1; ; i++ {
		v := heap.Pop(h).(int)
		if k == i {
			return v
		}
		for _, f := range factors {
			add := v * f
			if _, ok := m[add]; !ok {
				heap.Push(h, add)
				m[add] = true
			}
		}
	}
}

// 方法二： 动态规划
// 令dp[k] = result, 表示第k个数的结果为result
func getKthMagicNumber_2(k int) int {
	dp := make([]int, k + 1)
	dp[1] = 1
	// 初始化3个因子对应下标值
	x, y, z := 1, 1, 1
	for i := 2; i <= k; i++ {
		// 如何计算dp[i]
		// 分别取对应因子的下标对应的dp[x/y/z] 与因子相乘，取最小
		// 判断最后的结果与哪个下标对应的计算值相等，将其下标+1
		x1, y1, z1 := dp[x]*3, dp[y]*5, dp[z]*7
		dp[i] = min(min(x1, y1), z1)
		if dp[i] == x1 {
			x++
		}
		if dp[i] == y1 {
			y++
		}
		if dp[i] == z1 {
			z++
		}
	}
	return dp[k]
}

//leetcode submit region end(Prohibit modification and deletion)
