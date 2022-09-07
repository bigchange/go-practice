package leetcode

import (
	"fmt"
	"strings"
)

func reorderSpaces(text string) string {
	spaces := 0
	var words []string
	isStarted := false
	startPos := 0
	// 统计空格个数
	// 以及将其中的单词提出来
	for index, i := range text {
		if i != ' ' {
			if !isStarted {
				// 记录单词的开始位置
				isStarted = true
				startPos = index
			}
		} else {
			// 累积空格数
			spaces++
			// 遇到第一个空格，判断是否已经记录单词的开始位置
			if isStarted {
				words = append(words, text[startPos:index])
			}
			// 将记录单词的开始位置重置
			isStarted = false
		}
	}
	// 最后一个单词没有空格的时候, 看是否已经记录了单词的开始位置
	if isStarted {
		words = append(words, text[startPos:len(text)])
	}
	size := len(words)
	// 每个单词间的空格个数
	var spaceNumBetweenWord int
	// 边界case
	if size == 1 {
		spaceNumBetweenWord = 0
	} else {
		spaceNumBetweenWord = spaces / (size - 1)
	}
	// 剩余空格
	leftSpaces := spaces - (spaceNumBetweenWord * (size - 1))
	spacesIn := strings.Repeat(" ", spaceNumBetweenWord)
	sb := strings.Builder{}
	fmt.Println("words:", words)
	for i, word := range words {
		if i == 0 {
			sb.WriteString(word)
		} else {
			sb.WriteString(spacesIn)
			sb.WriteString(word)
		}
	}
	sb.WriteString(strings.Repeat(" ", leftSpaces))
	fmt.Println("spaces:", spaces, leftSpaces, spaceNumBetweenWord, size)
	retStr := sb.String()
	fmt.Println("ret:",fmt.Sprintf(",%v,", retStr))
	return retStr
}
