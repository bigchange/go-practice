package leetcode

import (
	"strings"
)

// 每行对齐分为一下情况：
// 1. 如果只有一个单词，左对齐
// 2. 如果是最后一行，左对齐
// 3. 其他情况： 根据单词数，判断个间隙应该补充的空格数量

// 思考： 如果判断第3种情况下单词数？
// 根据题意： 应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词
// 简单理解： 就假设单词之间就只存在一个空格， 从左到右依次累加每个单词的长度（包括单词本身之间最少是一个空格）得到总长度
// 判断总长度是否超过maxWidth
// 如果在累加某个位置（i）单词长度时超过了
// 那么之前的单词就可以组成一行，然后补充空格进去就可以了

// 每个间隙空格数的计算逻辑：
// 每行单词个数为： n ， 间隙数为： n - 1， 单词总长度： 每个单词长度累加
// 每行空格总数为: t = maxWidth - 单词总长度
// 每行每个间隙空格初始数为：m = t / (n-1)
// 空格余数： left = t % (n-1)
// 从左到右，前left个间隙的空格数为： m + 1， 剩余的间隙空格数为： m

func fullJustify(words []string, maxWidth int) []string {
	// 结果集合
	var ans []string
	n := len(words)
	// 从i位置为起点开始获取每行能有的最多单词数据
	// 返回：当前行单词集合，下次的起点位置，单词总长度
	getRows := func (i int) ([]string, int, int) {
		var rows []string
		// 贪心
		// 尽可能多的每一行有更多的单词，不过需要注意每个单词间需要至少一个空格
		// 如果超过maxWidth，开始计算每个单词之前间隔的数量，返回最终row字符串
		t := len(words[i])
		rows = append(rows, words[i])
		minSpace := 1
		for  i + 1  < n && t + minSpace + len(words[i + 1]) <= maxWidth {
			rows = append(rows, words[i + 1])
			t += minSpace + len(words[i + 1])
			i++
		}
		return rows, i + 1, t
	}
	start := 0
	for  {
		rows, nextPos, totalLength := getRows(start)
		// 左对齐处理
		// 表示数据已经都遍历完了，rows里面的数据就是最后一行了
		if nextPos == n {
			 s := strings.Join(rows, " ")
			 cl := len(s)
			 for i := cl; i < maxWidth;  i++ {
			 		s += " "
			 }
			 ans = append(ans,  s)
			 break
		} else {
			// 如果只有一个单词
			// 左对齐
			l := len(rows)
			if l == 1 {
				s := rows[0]
				cl := len(s)
				for i := cl; i < maxWidth;  i++ {
					s += " "
				}
				ans = append(ans,  s)
			} else {
				// 其他情况
				l := len(rows) // 单词数据
				// 间隙的个数
				spaceNums := l - 1
				// 单词总长度
				totalWordsLen := totalLength - spaceNums
				// 空格总个数
				totalSpaceNums := maxWidth - totalWordsLen
				// 每个间隙空字符串的个数
				eachSpaceItemNums :=  totalSpaceNums / spaceNums
				leftSpaceNums := totalSpaceNums % spaceNums
				s := ""
				for j := 0; j < l; j++ {
					s += rows[j]
					// 最后一个字符串结束
					if j == l - 1 {
						break
					}
					// 补充空格
					spaces := eachSpaceItemNums
					if j < leftSpaceNums {
						spaces = spaces + 1
					}
					for k := 0; k < spaces; k++ {
						s += " "
					}
				}
				ans = append(ans, s)
			}
		}
		start = nextPos
	}
	return ans
}

