package leetcode

/*
给定一个整形数组，求连续子数组的最大和（连续子数组乘积最大也是类似问题）
[1, 2, -1, 3, 4, -2]
1+2=3  // valid
3+4=7  // valid
1+2+-1+3+4=9  //valid, max
1+2+-10+3+4
1+2+3+4 //invalid

// 如果使用DP
// 1. 连续子数组最大和：那么状态转移方程为，dp[i] = max(dp[i-1]+ nums[i], nums[i])
// return max(dp[i]...)

// 2. 连续子数组乘积最大：思路：当前数乘以上一个最大值，当前值，当前数乘以上一个最小值，这三者比较，其中的最大值，就是我们要的最大值，那么状态转移方程为
// maxF[i] = max(maxF[i-1] * nums[i], max(nums[i], minF(i-1] * nums[i]))
// minF[i] = min[minF[i-1] * nums[i], min(nums[i], minF[i-1] * nums[i]))
// return max(maxF[i]...)
*/

func maxSubArray(nums []int) int {
	currSum := nums[0]
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if currSum <= 0 {
			currSum = nums[i]
		} else {
			currSum += nums[i]
		}
		if currSum > m {
			m = currSum
		}
	}
	return m
}
