package leetcode

// 当 k = 1，按1，2，3，4，..，n 排列即可
// 当 k = n - 1的时候
// 按如下排列： 1,n,2,n-1,3,n-2,4,n-3,,,,k,n-k-1
// 一般情况，可以将两者结合
// 前半部分差值都为k = 1， 后半部分按差值 k 逐渐减少
// 这样 1 到 k的差值均出现一次，满足条件
func constructArray(n int, k int) []int {
	ret := make([]int, n)
	// 前半部分： [1，2，3，4，， ~ n - k - 1] 相差都为1
	for i:= 1; i < n-k; i++ {
		ret[i - 1] = i
	}
	// 后半部分按 n-k,n,n-k-1,n-1,n-k-2,n-2,... 排列就可以了
	i,j := n-k,n
	cnt := 0
	for i <= j {
		ret[i - 1 + cnt] = i
		// 不能添加重复的进去
		if i != j {
			ret[i + cnt] = j
		}
		cnt++
		i++
		j--
	}
	return ret
}
