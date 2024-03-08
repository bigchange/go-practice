package leetcode
// 学校的自助午餐提供圆形和方形的三明治，分别用数字 0 和 1 表示。所有学生站在一个队列里，每个学生要么喜欢圆形的要么喜欢方形的。 餐厅里三明治的数量与学生
//的数量相同。所有三明治都放在一个 栈 里，每一轮：
//
//
// 如果队列最前面的学生 喜欢 栈顶的三明治，那么会 拿走它 并离开队列。
// 否则，这名学生会 放弃这个三明治 并回到队列的尾部。
//
//
// 这个过程会一直持续到队列里所有学生都不喜欢栈顶的三明治为止。
//
// 给你两个整数数组 students 和 sandwiches ，其中 sandwiches[i] 是栈里面第 i 个三明治的类型（i = 0 是栈的顶部）
//， students[j] 是初始队列里第 j 名学生对三明治的喜好（j = 0 是队列的最开始位置）。请你返回无法吃午餐的学生数量。
//
//
//
// 示例 1：
//
// 输入：students = [1,1,0,0], sandwiches = [0,1,0,1]
//输出：0
//解释：
//- 最前面的学生放弃最顶上的三明治，并回到队列的末尾，学生队列变为 students = [1,0,0,1]。
//- 最前面的学生放弃最顶上的三明治，并回到队列的末尾，学生队列变为 students = [0,0,1,1]。
//- 最前面的学生拿走最顶上的三明治，剩余学生队列为 students = [0,1,1]，三明治栈为 sandwiches = [1,0,1]。
//- 最前面的学生放弃最顶上的三明治，并回到队列的末尾，学生队列变为 students = [1,1,0]。
//- 最前面的学生拿走最顶上的三明治，剩余学生队列为 students = [1,0]，三明治栈为 sandwiches = [0,1]。
//- 最前面的学生放弃最顶上的三明治，并回到队列的末尾，学生队列变为 students = [0,1]。
//- 最前面的学生拿走最顶上的三明治，剩余学生队列为 students = [1]，三明治栈为 sandwiches = [1]。
//- 最前面的学生拿走最顶上的三明治，剩余学生队列为 students = []，三明治栈为 sandwiches = []。
//所以所有学生都有三明治吃。
//
//
// 示例 2：
//
// 输入：students = [1,1,1,0,0,1], sandwiches = [1,0,0,0,1,1]
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= students.length, sandwiches.length <= 100
// students.length == sandwiches.length
// sandwiches[i] 要么是 0 ，要么是 1 。
// students[i] 要么是 0 ，要么是 1 。
//
//
// Related Topics 栈 队列 数组 模拟 👍 82 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// P 入口
func P(students []int, sandwiches []int) int {
	size := len(sandwiches)
	var getSand func(int) int
	// startPos：sandwiches的开始位置
	// 返回当前还有多少同学在队列中
	getSand = func(startPos int) int {
		// 终止条件
		// 所有sandwiches都被拿走
		if startPos >= size {
			return 0
		}
		// 当前学生的数量
		studentLen := len(students)
		count := 0
		// 判断剩下的学生是否有能拿走队列sandwiches的
		for k := 0; k < studentLen; k++ {
			if students[k] != sandwiches[startPos] {
				count++
				// 放入队列尾部
				students = append(students, students[k])
			} else {
				// 得到新的排列同学队列
				students = students[k+1:]
				break
			}
		}
		// 如果剩余同学都无法拿走队列的sandwiches，该队列学生都无法拿到sandwiches
		// 返回队列长度
		if count == studentLen {
			return studentLen
		}
		// 新的学生对队列，开始下一个sandwiches的匹配
		return getSand(startPos + 1)
	}
	// 从第一个 sandwiches 开始匹配
	return getSand(0)
}
//leetcode submit region begin(Prohibit modification and deletion)
