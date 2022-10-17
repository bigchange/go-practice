package leetcode

//你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i] 是第 i 棵树上的水果 种类 。
//
// 你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：
//
//
// 你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
// 你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合篮子中的水果类型。每采摘一次，你将会向右移动到
//下一棵树，并继续采摘。
// 一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
//
//
// 给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。
//
//
//
// 示例 1：
//
//
//输入：fruits = [1,2,1]
//输出：3
//解释：可以采摘全部 3 棵树。
//
//
// 示例 2：
//
//
//输入：fruits = [0,1,2,2]
//输出：3
//解释：可以采摘 [1,2,2] 这三棵树。
//如果从第一棵树开始采摘，则只能采摘 [0,1] 这两棵树。
//
//
// 示例 3：
//
//
//输入：fruits = [1,2,3,2,2]
//输出：4
//解释：可以采摘 [2,3,2,2] 这四棵树。
//如果从第一棵树开始采摘，则只能采摘 [1,2] 这两棵树。
//
//
// 示例 4：
//
//
//输入：fruits = [3,3,3,1,2,1,1,2,3,3,4]
//输出：5
//解释：可以采摘 [1,2,1,1,2] 这五棵树。
//
//
//
//
// 提示：
//
//
// 1 <= fruits.length <= 10⁵
// 0 <= fruits[i] < fruits.length
//
//
// Related Topics 数组 哈希表 滑动窗口 👍 318 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func totalFruit(fruits []int) int {
	start := 0
	// 已经采摘的最多果树数量
	getMax := 0
	size := len(fruits)
	// 开始采摘位置start的最大边界
	for start < (size - getMax) {
		// 果树类型 -> 位置下标的映射
		m := make(map[int]int)
		// 遇到的一个和第二个不同类型果树的类型
		first, sec := -1, -1
		// 当前开始位置能拿到的最多数量
		canMax := 0
		// 从当前位置start开始	采摘
		for i := start; i < size; i++ {
			item := fruits[i]
			// 与两个篮子中的第一个果树类型相同
			if item == first {
				first = item
				// 更新该果树类型的最新位置下标
				m[first] = i
				// 结果+1
				canMax++
				continue
			}
			if sec == item {
				// 与两个篮子中的第二个果树类型相同
				sec = item
				// 保留第二个不同类型果树的最初位置，不更新成最新位置下标
				// 与第一个不同类型的果树更新不一样，注意区别
				if _, ok := m[sec]; !ok {
					m[sec] = i
				}
				// 结果+1
				canMax++
				continue
			}
			// 遇到了两个不同的果树
			// 此开始start位置的采摘提前结束，直接提出此循环
			if first != -1 && sec != -1 {
				break
			}
			// 保留遇到的第一个果树类型
			if first == -1 && item != first {
				first = item
				m[first] = i
				// +1
				canMax++
			} else if item != sec {
				// 保留遇到的第二个不同果树类型
				sec = item
				m[sec] = i
				// +1
				canMax++
			} else {
				// 中途遇到第三不同类型果树时：
				// 下一次开始采摘的位置，从本次采摘中遇到的第二个不同类型果树的最初位置开始
				// 开始下次采摘过程，提前退出
				start = m[sec]
				// 保存本次采摘数量到最终返回结果变量中
				getMax = max(getMax, canMax)
				break
			}
		}
		// 特殊情况：如果一直没遇到第三种类型果树
		// 同样需要计算此start位置开始采摘的结果并保存
		start = m[sec]
		getMax = max(getMax, canMax)
	}
	// 返回最终结果
	return getMax
}
//leetcode submit region end(Prohibit modification and deletion)

