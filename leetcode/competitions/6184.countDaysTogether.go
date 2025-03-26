/*
 * @Author: Jerry You
 * @CreatedDate: 2022/9/17
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/9/17 22:39
 */
package competitions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// CountDaysTogether 计算两个人在一段时间内重合的天数
// arriveAlice: Alice到达的日期
// leaveAlice: Alice离开的日期
// arriveBob: Bob到达的日期
// leaveBob: Bob离开的日期
// 返回值: Alice和Bob重合的天数
func CountDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	type d struct {
		N string
		D string
	}
	// 同一天到达，同一天离开
	// 特殊处理这种case: "01-01", "01-01", "01-01", "01-01"
	// 或者初始化排序的数组使用 a b a b的形式
	if arriveAlice == arriveBob && leaveAlice == leaveBob {
		return calDay(arriveAlice, leaveAlice)
	}
	var ds = []d{
		{
			N: "a",
			D: arriveAlice,
		},
		{
			N: "b",
			D: arriveBob,
		},
		{
			N: "a",
			D: leaveAlice,
		},
		{
			N: "b",
			D: leaveBob,
		},
	}
	// 按日期排序
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].D < ds[j].D
	})
	// 如果第0个和第1个时间为同一个人的
	if ds[0].N == ds[1].N {
		return 0
	} else {
		// 不然，重合时间为第1个和第2个
		// 计算天数即可
		return calDay(ds[1].D, ds[2].D)
	}
}
func calDay(start, end string) int {
	sum := 0
	fmt.Println(start, end)
	starts := strings.Split(start, "-")
	ends := strings.Split(end, "-")
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if starts[0] == ends[0] {
		return toInt(ends[1]) - toInt(starts[1]) + 1
	} else {
		sm := toInt(starts[0])
		sd := toInt(starts[1])
		em := toInt(ends[0])
		ed := toInt(ends[1])
		sum = days[sm-1] - sd + 1
		for i := sm; i < em-1; i++ {
			sum += days[i]
		}
		sum += ed
	}
	return sum
}

func toInt(s string) int {
	if strings.HasPrefix(s, "0") {
		v, _ := strconv.Atoi(s[1:])
		return v
	} else {
		v, _ := strconv.Atoi(s[0:])
		return v
	}
}
