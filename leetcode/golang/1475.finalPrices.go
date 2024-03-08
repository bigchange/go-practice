package leetcode

// 入口
func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		prices[i] = discountPrice(prices, i)
	}
	return prices
}

// 直接遍历
// 当前位置pos 之后比其val小的第一个val为其折扣价
func discountPrice(prices []int, pos int) int {
	size := len(prices)
	i := pos
	for i = pos + 1; i < size; i++ {
		// 找到对应折扣价,返回折扣后的价格
		if  prices[i] <= prices[pos] {
			return prices[pos] -  prices[i]
		}
	}
	// 返回原价
	return prices[pos]
}

// 方法二：单调栈
// 维护当前位置右边第一个更小的值列表
// 栈底到栈顶元素单调递增
// 使用slice维护一个栈
func finalPrices_2(prices []int) []int {
	ans := make([]int, len(prices))
	stack := []int{0}
	// 倒序遍历
	for i := len(prices) - 1; i >= 0; i-- {
		p := prices[i]
		// 栈顶比当前价格大，出栈
		// 思路：当前位置前面的元素右边第一个更小的值，是当前位置
		// 栈内比当前位置更大的可以直接丢弃了
		for len(stack) > 1 && stack[len(stack) - 1] > p {
			stack = stack[:len(stack) - 1]
		}
		// 计算结果
		ans[i] = p - stack[len(stack) - 1]
		// 当前位置值入栈
		stack = append(stack, p)
	}
	return ans
}
