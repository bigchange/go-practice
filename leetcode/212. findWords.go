package leetcode

import "fmt"


// corner case

// [[a]]
// [a]
//
// [["o","a","b","n"],["o","t","a","e"],["a","h","k","r"],["a","f","l","v"]]
// ["oa","oaa"]

func findWords(board [][]byte, words []string) []string {
  r := len(board)
  c := len(board[0])
  var rets []string
  total := make([]byte, 26)
  // 定义搜索的方向
  directions := [][]int{{0,1},{0, -1},{1, 0},{-1, 0}}
	isMatched := false

	// 预处理: 统计board中每个字符的个数
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			total[board[i][j] - 'a']++
		}
	}
	// 剪枝优化
	// 比较字符串每个字符的个数
	// 与board中每个字符的个数
	// 如果board中存在某个字符个数小于匹配字符串中的某个字符个数
	// 那该字符串肯定匹配不上
	checkValid := func(word string) bool {
		totalInW := make([]byte, 26)
		for _, w := range word {
			totalInW[w - 'a']++
		}
		for i := 0; i < 26; i++ {
			if total[i] < totalInW[i] {
				return false
			}
		}
		return true
	}
  var dfs func([][]byte, string, int, int, int) bool
  dfs = func(board [][]byte, str string, i, j, l int) bool {
  	fmt.Printf("searching index:%v,%v, l:%v\n", i, j, l)
  	// 终止条件
  	// 当递归到超过匹配字符串的长度时，匹配结束
  	if l == len(str)  {
			// fmt.Printf("return string:%v,%v, l:%v\n", i, j, l)
  		if !isMatched {
				rets = append(rets, str)
				isMatched = true
			}
  		return isMatched
		}
		// 如果与当前字符串相等
		// 可以继续搜索下去
		currentVal := board[i][j]
		if currentVal != str[l] {
			// fmt.Printf("%v,not equal to: %v,\n", currentVal, str[l])
			return false
		} else {
			//  如果下一个字符比较已经出现超过匹配字符串的长度
			//  那么也将终止
			if l + 1 == len(str) {
				if !isMatched {
					rets = append(rets, str)
					isMatched = true
				}
				return isMatched
			}
		}
		// fmt.Printf("matching \n")
		board[i][j] = '0'
		// 继续按照四个方向进一步搜索
		for d := 0; d < len(directions); d++ {
			nx := i + directions[d][0]
			ny := j + directions[d][1]
			// 按某个方向继续下去的条件
			// 1. 下标位置需要满足的边界
			// 2. 访问过的位置不能重复（通过将访问过位置的元素置为一个特殊值来界定）
			if nx >= 0 && nx < r && ny >= 0 && ny < c && board[nx][ny] != '0' {
				dfs(board, str, nx, ny, l + 1)
			}
		}
		// 四个方向都尝试过后
		// 递归回溯，需要将重置的值恢复
		board[i][j] = currentVal
		// 返回当前递归的匹配结果
		return isMatched
	}

	for _, w := range words {
		isMatched = false
		if !checkValid(w) {
			continue
		}
		for i := 0; i < r && !isMatched; i++ {
			for j := 0; j < c && !isMatched; j++ {
				isMatched = dfs(board, w, i, j, 0)
			}
		}
	}
	return rets
}
