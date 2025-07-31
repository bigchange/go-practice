package leetcode

// 如何在两个正序的数组中找到其中位数
// 思路转化： 找到第k小的数， 开始下标位置：k = (m + n ) / 2 或者 (m + n ) / 2 + 1
// 由于两个数组长度已知（m，n）
// 那么中位数的位置就确定：
// if (m+n) % 2 == 0 (总元素偶数个）
// k1 = (m + n ) / 2 或 k = (m + n ) / 2 + 1
// 那么此时需要找到：一个第k1小的数 和 一个第k2小的数（每次找的逻辑一样，实现可通过分开找两次，具体实现可自定义）
// 比如： m + n = 4, 那么需要找第 k1=2小的数和 第k2=3小的数
// else （总元素奇数个）
// k = (m + n) / 2 + 1
// 那么此时需要找：一个第k小的数即可
// 比如： m + n = 3,  那么需要找第k=2小的数
// 通过找到中位数的位置，那么中位数就间接得到了

// 二分查找条件： 通过一次比较来排除最多可能是k/2 个不可能是第 k 小的数，查找范围缩小了一半。同时，我们将在排除后的新数组上继续进行二分查找，并且根据我们排除数的个数，减少 k 的值
// 1. if A[start + k/2 - 1] <= B[start + k/2 - 1]， 那么可以排除A[start] ~ A[start + k/2-1] 的数
// 2  if A[start + k/2 - 1] <= B[start + k/2 - 1]， 那么可以排除 B[start] ~ A[start + k/2-1] 的数
// (如此循环）在剩余A数组中和剩余B数组中继续按照上面的逻辑查找第k小的数， 此时： k = k - 每一轮已排除数组A或B中元素的个数；

// 注意说明：
// 下标位置：k/2 - 1 不是数组中的绝对位置，是相对于每轮排除一定元素后，重新按0开始计数下标后的相对下标位置
// 比如，A中已经被排除了3个元素，此时 若 k = 2， 有效元素的开始位置 start = 3（0,1,2下标的元素被排除了)，那么此时
// 下标位置：k/2 - 1 指定的位置在A数组中是： pos = start + k / 2 - 1 = 3 + 2 / 2 - 1 =  3， 对应的元素就是 A[3]

// 需要处理的边界条件：
// 1. 如果其中某一个数组中位置越界（都被排除了），那么在另一个数组中直接返回第k小的数
// 2. 如果 k=1, 那么第k小的数就是 A和B数组中第一个有效元素的最小值
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	pos := (m + n) / 2
	if (m+n)%2 == 0 {
		r1 := getKMinNum(nums1, nums2, pos)
		r2 := getKMinNum(nums1, nums2, pos+1)
		return (r1 + r2) / 2.0
	}
	r := getKMinNum(nums1, nums2, pos+1)
	return r
}
func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
func getKMinNum(nums1 []int, nums2 []int, k int) float64 {
	index1, index2 := 0, 0
	for {
		if index2 >= len(nums2) {
			return float64(nums1[index1+k-1])
		}
		if index1 >= len(nums1) {
			return float64(nums2[index2+k-1])
		}
		if k == 1 {
			return float64(min(nums2[index2], nums1[index1]))
		}
		half := k / 2
		// 处理一下
		pos1 := min(index1+half, len(nums1)) - 1
		pos2 := min(index2+half, len(nums2)) - 1
		if nums1[pos1] <= nums2[pos2] {
			k = k - (pos1 - index1 + 1)
			index1 = pos1 + 1
		} else {
			k = k - (pos2 - index2 + 1)
			index2 = pos2 + 1
		}
	}
}
