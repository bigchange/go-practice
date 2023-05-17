package backtracking

// 回溯算法：该算法在搜索解空间时会采用“尝试”与“回退”的策略

// 可通过解决数独问题，来深入理解回溯， 代码位置：(leetcode/java/problems/SudokuSolver37.java)

// 尝试和回退策略
// 在相关约束条件上的： 剪枝

// 常用术语： 解、状态、约束条件等术语是通用的，适用于回溯算法、动态规划、贪心算法等。

// 典型问题
// 搜索问题： 全排列， 子集合，汉诺塔
// 约束满足: N皇后，图着色问题
// 组合优化： 0-1背包， 旅行商问题，最大团问题
//

/* 回溯算法框架
// 设 state 为问题的当前状态，choices 表示当前状态下可以做出的选择
func backtrack(state *State, choices []Choice, res *[]State) {
	// 判断是否为解
	if isSolution(state) {
		// 记录解
		recordSolution(state, res)
		return
	}
	// 遍历所有选择
	for _, choice := range choices {
		// 剪枝：判断选择是否合法
		if isValid(state, choice) {
			// 尝试：做出选择，更新状态
			makeChoice(state, choice)
			backtrack(state, choices, res)
			// 回退：撤销选择，恢复到之前的状态
			undoChoice(state, choice)
		}
	}
}
*/
