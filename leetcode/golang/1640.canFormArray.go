package leetcode

//给你一个整数数组 arr ，数组中的每个整数 互不相同 。另有一个由整数数组构成的数组 pieces，其中的整数也 互不相同 。请你以 任意顺序 连接
//pieces 中的数组以形成 arr 。但是，不允许 对每个数组 pieces[i] 中的整数重新排序。
//
// 如果可以连接 pieces 中的数组形成 arr ，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：arr = [15,88], pieces = [[88],[15]]
//输出：true
//解释：依次连接 [15] 和 [88]
//
//
// 示例 2：
//
//
//输入：arr = [49,18,16], pieces = [[16,18,49]]
//输出：false
//解释：即便数字相符，也不能重新排列 pieces[0]
//
//
// 示例 3：
//
//
//输入：arr = [91,4,64,78], pieces = [[78],[4,64],[91]]
//输出：true
//解释：依次连接 [91]、[4,64] 和 [78]
//
//
//
// 提示：
//
//
// 1 <= pieces.length <= arr.length <= 100
// sum(pieces[i].length) == arr.length
// 1 <= pieces[i].length <= arr.length
// 1 <= arr[i], pieces[i][j] <= 100
// arr 中的整数 互不相同
// pieces 中的整数 互不相同（也就是说，如果将 pieces 扁平化成一维数组，数组中的所有整数互不相同）
//
// Related Topics 数组 哈希表 👍 88 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func canFormArray(arr []int, pieces [][]int) bool {
	visited := make(map[int]bool)
	size := len(arr)
	canFormCheckFunc := func(targetIndex int) int {
		for index, i := range pieces {
			if targetIndex >= size {
				return size
			}
			if _, ok := visited[index];ok {
				continue
			}
			tmp := targetIndex
			shouldMatches := len(i)
			// 判读该piece是否和arr对应位置依次连续相等
			for _, p := range i {
				if arr[targetIndex] == p {
					targetIndex++
					continue
				} else {
					break
				}
			}
			// 如果没有一个相等
			if targetIndex == tmp {
				continue
			}
			// 如果相等的数不满piece数组长度，直接返回特殊值，表示最终结果无法找到
			// 前提：pieces 中的整数 互不相同（也就是说，如果将 pieces 扁平化成一维数组，数组中的所有整数互不相同）
			// 如果存在相同的，就不能直接返回，需要进一步去匹配其他的piece
			if targetIndex - tmp != shouldMatches {
				return -1
			}
			// 标记已经使用了的数组
			visited[index] = true
		}
		// 返回已经匹配到了arr的哪个下标位置
		return targetIndex
	}
	// 开始位置
	start := 0
	for range  arr {
		// 返回已匹配到的arr的位置
		// 继续往下匹配
		start = canFormCheckFunc(start)
		// 提前结束
		if start == -1 {
			return false
		}
		// 匹配到末尾，可结束
		if start >= size {
			return true
		}
	}
	// 最后返回false
	return false

}
//leetcode submit region end(Prohibit modification and deletion)

