package leetcode

import (
	"math"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	// x 在最左边
	if  x <= arr[0] {
		return arr[:k]
	}
	size := len(arr)
	// x 在最右边
	if  x > arr[size -1] {
		return  arr[size - k:]
	}
	// x 在中间
	start, end := -1, -1
	for  i := 1; i < size; i++ {
		// 定位距离x最进的两个位置，start，end
		if  arr[i] >= x {
			start = i - 1
			end = i
			break
		}
	}
	i := start
	j := end
	var res  = make([]int, k)
	cnt := 0
	for i >= 0 || j < size {
		y, z := -1, -1
		// 找到K个就可以退出
		if cnt == k {
			break
		}
		pos := 0
		// 双指针移动
		// start向最左移动，end向右移动
		// 比较哪个距离x更近，取距离近的一边，同时移动该边的指针
	  if i >= 0 {
	  	y = distance(arr[i], x)
		}
		if j < size {
			z = distance(arr[j], x)
		}
		if  y != -1 && z != -1 {
			if y <= z {
				pos = i
				i--
			} else {
				pos = j
				j++
			}
		} else if y == -1 {
			// 左边越界，没有元素了，那就取右边
			pos = j
			j++
		} else {
			// 右边越界，没有元素了，那就取左边
			pos = i
			i--
		}
		res[cnt] = arr[pos]
		cnt++
	}
	// 最后排一下顺序
	sort.Ints(res)
	return  res
}

func distance(x, y int) int {
	return int(math.Abs(float64(x) - float64(y)))
}