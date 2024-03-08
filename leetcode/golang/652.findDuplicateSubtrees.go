package leetcode

import (
	"strconv"
	"strings"
)

// 入口
// 深度搜索
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var sameTrees []*TreeNode
	keyMap := make(map[string]int)
	var dfs func(r *TreeNode) string
	// 返回不同子树对应的标识
	dfs = func(r *TreeNode) string  {
		if r == nil {
			return ""
		}
		sb := strings.Builder{}
		sb.WriteString(strconv.Itoa(r.Val))

		ls := dfs(r.Left)
		sb.WriteString("_"+ ls)
		rs := dfs(r.Right)
		sb.WriteString("_"+ rs)
		retKey := sb.String()
		// 某一个子树对应的标识
		keyMap[retKey]++
		// 如果累积值到了2，那么就存在相同的子树，添加到结果集合中
		if keyMap[retKey] == 2 {
			sameTrees = append(sameTrees, r)
		}
		return retKey
	}
	dfs(root)
	return sameTrees

}


