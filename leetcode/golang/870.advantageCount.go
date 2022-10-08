package leetcode

import "sort"

//给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的
//数目来描述。
//
// 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
//输出：[2,11,7,15]
//
//
// 示例 2：
//
//
//输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
//输出：[24,32,8,12]
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length <= 10⁵
// nums2.length == nums1.length
// 0 <= nums1[i], nums2[i] <= 10⁹
//
//
// Related Topics 贪心 数组 双指针 排序 👍 280 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 方法一：排序
// 依次给nums1，nums2按递增顺序排序
// 遍历排好序的nums2，以此来寻找每个num应该与nums1中的哪个元素绑定
// 对于一个nums2中的数：
// 在排好序的nums1中从左到右依次找到第一个比其大的数据即可
// 同时，将nums1中对应位置的元素标记为已绑定
// 没有找到更大的元素时，随便找一个nums1中未被绑定的元素与之绑定即可
func advantageCount(nums1 []int, nums2 []int) []int {
	type Item struct {
		Num int
		// 记录其排序前初始位置下标
		Index int
	}
	var items = make([]Item, len(nums2))
	for i := 0; i < len(nums2); i++ {
		items[i] = Item{
			Num: nums2[i],
			Index: i,
		}
	}
	// 排序
	sort.Ints(nums1)
	sort.Slice(items, func(i, j int) bool {
		return items[i].Num < items[j].Num
	})
	j, k := 0, 0
	flags := make([]int, len(nums1))
	for _, item := range items {
		// 找更大的元素的位置下标k
		for k < len(nums1) && nums1[k] <= item.Num {
			k++
		}
		// 如果k超过最大下标
		if  k >= len(nums1) {
			// 随便找一个未被标记的即可
			for j < len(nums1) && flags[j] != 0 {
				j++
			}
			// 复用nums2，使用nums2保存最后的返回结果
			nums2[item.Index] = nums1[j]
			// 下标往后移动
			j++
		} else {
			nums2[item.Index] = nums1[k]
			// 同时标记该位置已被绑定
			flags[k] = 1
			// 下标往后移动
			k++
		}
	}
	return nums2
}

//leetcode submit region end(Prohibit modification and deletion)
