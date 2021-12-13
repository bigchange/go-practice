package leetcode

func interchangeableRectangles(rectangles [][]int) int64 {
	l := len(rectangles)
	d := func(x []int) float64 {
		x1 := float64(x[0]) / float64(x[1])
		return  x1
	}
	cMap := make(map[float64][]int)
	var ans int64 = 0
	for i := 0; i < l; i++ {
		v := d(rectangles[i])
		if _, ok := cMap[v]; !ok {
			var vals []int
			vals = append(vals, i)
			cMap[v] = vals
		} else {
			cMap[v] = append(cMap[v], i)
		}
	}
	for _, v := range cMap {
		for  i := 1; i < len(v); i++ {
			ans += int64(i)
		}
	}
	return ans
}
