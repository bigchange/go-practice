package leetcode

import (
	"fmt"
	"sort"
)

func trimMean(arr []int) float64 {
	ret := 0.0
	sum := 0
	size := len(arr)
	fmt.Println("arr size", size)
	delItem := int(float64(size) * 0.05)
	sort.Ints(arr)
	fmt.Println("arr", arr)
	arr = arr[0:size - delItem]
	fmt.Println("arr 1", arr)
	arr = arr[delItem:]
	fmt.Println("arr 2", arr)
	for _, i := range arr {
		sum += i
	}
	ret = float64(sum) / float64(len(arr))
	return ret
}
