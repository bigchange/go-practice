/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/9
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/9 21:15
 */
package leetcode

import "strings"

func minOperations(logs []string) int {
	step := 0
	for _, log := range logs {
		if strings.HasPrefix(log,"../") {
			step--
			if step < 0 {
				step = 0
			}
		} else if strings.HasPrefix(log, "./") {
			// do nothing
		} else {
			step++
		}
	}
	return step
}