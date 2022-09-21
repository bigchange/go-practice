package leetcode

//对于某些非负整数 k ，如果交换 s1 中两个字母的位置恰好 k 次，能够使结果字符串等于 s2 ，则认为字符串 s1 和 s2 的 相似度为 k 。
//
// 给你两个字母异位词 s1 和 s2 ，返回 s1 和 s2 的相似度 k 的最小值。
//
//
//
// 示例 1：
//
//
//输入：s1 = "ab", s2 = "ba"
//输出：1
//
//
// 示例 2：
//
//
//输入：s1 = "abc", s2 = "bca"
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= s1.length <= 20
// s2.length == s1.length
// s1 和 s2 只包含集合 {'a', 'b', 'c', 'd', 'e', 'f'} 中的小写字母
// s2 是 s1 的一个字母异位词
//
// Related Topics 广度优先搜索 字符串 👍 167 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

// 广度搜索 + 剪枝
// 超时cases：
// "baababbaaabbabaaabaa","abaabaabababaabababa"
func kSimilarity(s1 string, s2 string) int {
	// 初始化返回结果：minK
	minK := len(s1)
	var arr []int32
	var targets []int32
	for _, i := range s1 {
		arr = append(arr, i)
	}
	for _, i := range s2 {
		targets = append(targets, i)
	}
	// 递归搜索
	// 参数： 当前位置index， 已经交换的次数num
	var bfs func(int, int)
	bfs = func(index int, num int) {
		// 如果交换到了最后一位，计算最小num保存
		if index == len(arr) {
			minK = min(num, minK)
			return
		}
		// 剪枝： 如果交互次数已经大于目前得到的最小交互次数，直接返回
		// 此剪枝效果明显，不加极易超时
		if num > minK {
			return
		}
		// 如果当前位置相等，继续递归下一个位置
		// 交互次数num不变
		if arr[index] == targets[index] {
			bfs(index+1, num)
			return
		}
		//从当前位置后一位开始，在s1中往后找与s2目标位置相同元素，两者交换位置
		for i := index + 1; i < len(arr); i++ {
			if arr[i] != targets[index] {
				continue
			}
			arr[index], arr[i] = arr[i], arr[index]
			// 继续递归下一个位置， 交换次数 + 1
			bfs(index + 1, num + 1)
			// s1中与s2目标位置相等的元素，应该不止一个，需要都搜索到
			// 此时需要将已经交换的元素复位，恢复现场，才能继续需要下一个交换的位置
			arr[i], arr[index] = arr[index], arr[i]
		}

	}
	// 从第一个位置开始搜
	bfs(0, 0)
	return minK
}

//leetcode submit region end(Prohibit modification and deletion)

