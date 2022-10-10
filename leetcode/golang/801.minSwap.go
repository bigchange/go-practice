/*
 * @Author: Jerry You
 * @CreatedDate: 2022/10/10
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/10/10 19:02
 */
package leetcode
//我们有两个长度相等且不为空的整型数组 nums1 和 nums2 。在一次操作中，我们可以交换 nums1[i] 和 nums2[i]的元素。
//
//
// 例如，如果 nums1 = [1,2,3,8] ， nums2 =[5,6,7,4] ，你可以交换 i = 3 处的元素，得到 nums1 =[1,2,3
//,4] 和 nums2 =[5,6,7,8] 。
//
//
// 返回 使 nums1 和 nums2 严格递增 所需操作的最小次数 。
//
// 数组 arr 严格递增 且 arr[0] < arr[1] < arr[2] < ... < arr[arr.length - 1] 。
//
// 注意：
//
//
// 用例保证可以实现操作。
//
//
//
//
// 示例 1:
//
//
//输入: nums1 = [1,3,5,4], nums2 = [1,2,3,7]
//输出: 1
//解释:
//交换 A[3] 和 B[3] 后，两个数组如下:
//A = [1, 3, 5, 7] ， B = [1, 2, 3, 4]
//两个数组均为严格递增的。
//
// 示例 2:
//
//
//输入: nums1 = [0,3,5,8,9], nums2 = [2,1,4,6,9]
//输出: 1
//
//
//
//
// 提示:
//
//
// 2 <= nums1.length <= 10⁵
// nums2.length == nums1.length
// 0 <= nums1[i], nums2[i] <= 2 * 10⁵
//
//
// Related Topics 数组 动态规划 👍 374 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	// 定义 f[i][j] 为考虑下标范围为 [0, i] 的元素，
	// 且位置 i 的交换状态为 j 时（其中 j = 0 为不交换，j = 1为交换）
	// 两数组满足严格递增的最小交换次数
	f := make([][]int, n)
	// 初始化
	f[0] = make([]int, 2)
	f[0][1] = 1
	for i := 1; i < n; i++ {
		arr := make([]int, 2)
		arr[0] = n + 10
		arr[1] = arr[0]
		f[i] = arr
	}
	// 开始状态转移计算
	for i := 1; i < n; i++ {
		// 即顺序位满足要求，此时要么当前位置 i 和前一位置 i−1 都不交换，要么同时发生交换
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			f[i][0] = f[i-1][0]
			f[i][1] = f[i-1][1] + 1
		}
		// 即交叉位满足要求，此时当前位置 i 和前一位置 i−1 只能有其一发生交换，此时有（分别对应「前一位置交换」和「当前位置交换」）
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			f[i][0] = min(f[i][0], f[i-1][1])
			f[i][1] = min(f[i][1], f[i-1][0]+1)
		}
	}
	// 最终结果
	return min(f[n-1][0], f[n-1][1])
}
//leetcode submit region end(Prohibit modification and deletion)
