/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/25
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/25 09:14
 */
package leetcode

import "strconv"

//我们称一个数 X 为好数, 如果它的每位数字逐个地被旋转 180 度后，我们仍可以得到一个有效的，且和 X 不同的数。要求每位数字都要被旋转。
//
// 如果一个数的每位数字被旋转以后仍然还是一个数字， 则这个数是有效的。0, 1, 和 8 被旋转后仍然是它们自己；2 和 5 可以互相旋转成对方（在这种情况
//下，它们以不同的方向旋转，换句话说，2 和 5 互为镜像）；6 和 9 同理，除了这些以外其他的数字旋转以后都不再是有效的数字。
//
// 现在我们有一个正整数 N, 计算从 1 到 N 中有多少个数 X 是好数？
//
//
//
// 示例：
//
// 输入: 10
//输出: 4
//解释:
//在[1, 10]中有四个好数： 2, 5, 6, 9。
//注意 1 和 10 不是好数, 因为他们在旋转之后不变。
//
//
//
//
// 提示：
//
//
// N 的取值范围是 [1, 10000]。
//
//
// Related Topics 数学 动态规划 👍 134 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 枚举每一个数
func rotatedDigits(n int) int {
	// 检查数据需要满足：
	// 1. 没有出现3，4，7
	// 2. 至少出现2，5，6，9
	checkInt := func(t int) (bool, bool) {
		// 是否满足1
		condition1 := true
		// 是否满足2
		condition2 := false
		for t != 0 {
			r := t % 10
			if r == 3 || r == 4 || r == 7 {
				condition1 = false
			}
			if r == 2 || r == 5 || r == 6 || r == 9 {
				condition2 = true
			}
			t = t / 10
		}
		return condition1, condition2
	}
	ret := 0
	for i := 1; i <= n; i++ {
		c1, c2 := checkInt(i)
		if c1 && c2 {
			ret++
		}
	}
	return ret
}

// 记忆化搜索: 可用来加深dfs的原理，写出更好的dfs代码
// dfs(pos, bound, diff) 分析：
// pos: 只从第 pos 位开始考虑。这里数的最高位为第 0 位，最低位为第 len−1 位
// bound参数：从第 0 位到第 pos−1 位的数是否「贴着」n，和n对应的位置数一样, 两种状态值：true or false
// diff参数：表示其是否为好数状态， 两种状态值：true or false
// 最终结果返回： dfs(0,True, false)
func rotatedDigits_2(n int) int {
	var check = [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}
	digits := strconv.Itoa(n)
	// n最大为10000，所以最高位数为5位
	// 第0维：位数
	// 第1维：bound状态
	// 第2维：diff状态
	memo := [5][2][2]int{}
	for i := 0; i < 5; i++ {
		memo[i] = [2][2]int{{-1, -1}, {-1, -1}}
	}
	var dfs func(int, bool, bool) int
	dfs = func(pos int, bound, diff bool) (res int) {
		// 所有位置都确定了，是否好数，根据diff状态判断即可
		if pos == len(digits) {
			return bool2int(diff)
		}
		ptr := &memo[pos][bool2int(bound)][bool2int(diff)]
		// 如果不是默认值-1，代表之前存在递归已经计算好了，直接返回计算结果
		if *ptr != -1 {
			return *ptr
		}
		// 默认：0 ~ 9个数字可选择
		lim := 9
		if bound {
			// 当前位置pos可能的值不能超过n所在该位置的数，不然就超过n了
			lim = int(digits[pos] - '0')
		}
		// 枚举可能
		for i := 0; i <= lim; i++ {
			if check[i] != -1 {
				res += dfs(pos+1, bound && i == int(digits[pos]-'0'), diff || check[i] == 1)
			}
		}
		*ptr = res
		return
	}
	return dfs(0, true, false)
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
//leetcode submit region end(Prohibit modification and deletion)
