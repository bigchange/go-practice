/*
 * @Author: Jerry You
 * @CreatedDate: 2023/3/15
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/3/15 13:02
 */

// Package exercism
package exercism

import (
	"fmt"
	"regexp"
	"strings"
)


func IsValidLine(text string) bool {
	// re := regexp.MustCompile(`^[TRC|DBG|INF|WRN|ERR|FTL]$ \S*`)
	// return re.MatchString(text)
	prefixes := []string{"TRC", "DBG", "INF", "WRN", "ERR", "FTL"}
	for _, p := range prefixes {
		if strings.HasPrefix(text, fmt.Sprintf("[%v]", p)) {
			return true
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile("\\<[~|*|\\-|=]*\\>").Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	c := 0
	reg  := regexp.MustCompile(`".*password.*"`)
	for _, line := range lines {
		lowerLine := strings.ToLower(line)
		if reg.MatchString(lowerLine) {
			c++
		}
	}
	return c
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	findUser := "User"
	var result = make([]string, len(lines))
	for i, line := range lines {
		startIndex := strings.Index(line, findUser)
		if startIndex == -1 {
			result[i] = line
		} else {
			startIndex = startIndex + len(findUser)
			if line[startIndex] != ' ' {
				result[i] = line
				continue
			}
			nameS, nameE := 0, 0
			for j, l := range line {
				if j < startIndex {
					continue
				}
				if  l == ' ' && nameS != 0 {
					nameE = j
					break
				}
				if  l != ' ' && nameS == 0 {
					nameS = j
				}
			}
			name := line[nameS: nameE]
			result[i] = fmt.Sprintf("[USR] %v %v", name, line)
		}
	}
	return result
}