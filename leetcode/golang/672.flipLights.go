/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/15
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/15 20:56
 */
package leetcode

import (
	"strconv"
	"strings"
)

// 方法二： 找规律
func flipLights_2 (n int, presses int) int {
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	} else if n == 2  {
		if presses == 1 {
			return 3
		} else {
			return 4
		}
	} else {
		if presses == 1 {
			return 4
		} else if presses == 2 {
			return 7
		} else {
			return 8
		}
	}
	return 0
}

// 方法一： 深度搜索（超出时间限制）
func flipLights(n int, presses int) int {
	lights := make([]bool, n + 1)
	for i := 1; i < n + 1; i++ {
		lights[i] = true
	}
	countMap := make(map[string]bool)
  switchCount := 4
  // 深度搜索
	var dfs func(lights []bool, press int)
	dfs = func(lights []bool, press int)  {
		// 按压开关次数到了
		if press == 0 {
			stat := getLightStat(lights)
			_, ok := countMap[stat]
			if !ok {
				countMap[stat] = true
			}
			return
		}
		originLight := make([]bool, len(lights))
		copy(originLight, lights)
		for i := 0; i < switchCount; i++ {
			lights = originLight
			switchPress(lights, i)
			dfs(lights, press - 1)
		}
	}
	dfs(lights, presses)
	return  len(countMap)
}

func getLightStat(lights []bool) string {
	sb := strings.Builder{}
	for i := 1; i < len(lights); i++ {
		sb.WriteString(strconv.FormatBool(lights[i]))
	}
	return sb.String()
}

func switchPress(lights []bool, i int) {
	k := 1
	step := 1
	size := len(lights)
	switch i {
	case 0:
	case 1:
		k = 2
		step = 2
	case 2,3:
		step = 2
	default:
		k = size
	}
	for k < size {
		lights[k] = !lights[k]
		if i == 3 {
			k = k + 1
			k = 3 * k + 1
		} else {
			k = k + step
		}
	}
}
