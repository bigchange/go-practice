package backtracking

// 全排列问题

// 给定数组，求数组元素全排列的可能

/* 全排列 I */
func permutationsI(nums []int) [][]int {
	// 排列结果存储
	res := make([][]int, 0)
	// 当前选择了的元素
	state := make([]int, 0)
	// 哪个元素被已经选过了，一个元素只能被选中一次
	selected := make([]bool, len(nums))
	backtrackI(&state, &nums, &selected, &res)
	return res
}
/* 全排列 II */
func permutationsII(nums []int) [][]int {
	// 排列结果存储
	res := make([][]int, 0)
	// 当前选择了的元素
	state := make([]int, 0)
	// 哪个元素被已经选过了，一个元素只能被选中一次
	selected := make([]bool, len(nums))
	backtrackII(state, nums, selected, &res)
	return res
}

/* 回溯算法：全排列 I */
// selected 和choices 应该可以不用指针类型传参吧？ 参考：backtrackII
func backtrackI(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	// 当状态长度等于元素数量时，记录解
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}
	// 遍历所有选择
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		// 剪枝：不允许重复选择元素 且 不允许重复选择相等元素
		if !(*selected)[i] {
			// 尝试：做出选择，更新状态
			(*selected)[i] = true
			*state = append(*state, choice)
			// 进行下一轮选择
			backtrackI(state, choices, selected, res)
			// 回退：撤销选择，恢复到之前的状态
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

/* 回溯算法：全排列 II */
// state，selected，choices 不用指针类型的参数
func backtrackII(state []int, choices []int, selected []bool, res *[][]int) {
	// 当状态长度等于元素数量时，记录解
	if len(state) == len(choices) {
		newState := append([]int{}, state...)
		*res = append(*res, newState)
	}
	// 遍历所有选择
	for i := 0; i < len(choices); i++ {
		choice := (choices)[i]
		// 剪枝：不允许重复选择元素 且 不允许重复选择相等元素
		if !(selected)[i] {
			// 尝试：做出选择，更新状态
			(selected)[i] = true
			state = append(state, choice)
			// 进行下一轮选择
			backtrackII(state, choices, selected, res)
			// 回退：撤销选择，恢复到之前的状态
			(selected)[i] = false
			state = (state)[:len(state)-1]
		}
	}
}
