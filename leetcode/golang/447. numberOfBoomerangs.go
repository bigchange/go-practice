package leetcode

// 暴力枚举 + hash表
// 每一轮固定位置i， 分别计算，其他节点与i的距离，将结果map中
// map: key：dis距离， val: 与i有相同距离的节点个数
// 然后通过遍历map中所有dis的key， 就可以得到不同dis下（ i， j， k) 的组合
// 每个dis下： （j,k) 元组的个数可以通过组合的方式计算得到
func numberOfBoomerangs(points [][]int) int {
	// 两个点之间的距离
	distance := func(x, y []int) int {
		x1 := x[0] - y[0]
		y1 := x[1] - y[1]
		return  x1 * x1 + y1 * y1
	}
	ans := 0
	for i, p1 := range  points {
		distanceMap := make(map[int]int)
		for j, p2 := range points {
			if j != i {
				dis := distance(p1, p2)
				distanceMap[dis]++
			}
		}
		// 当前i固定
		// l为与i有这相同距离的节点个数
		// 通过排列组合计算得出（j，k）的个数
		for _, l := range distanceMap {
			ans += l * (l - 1)
		}
	}
	return ans
}