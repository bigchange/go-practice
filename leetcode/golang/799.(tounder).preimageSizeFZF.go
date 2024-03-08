/*
 * @Author: Jerry You
 * @CreatedDate: 2022/8/28
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/8/28 18:44
 */
package leetcode

import "sort"

func preimageSizeFZF(k int) int {
 return nx(k +1) - nx(k)
}

func zeta(n int) (res int) {
	for n > 0 {
		n /= 5
		res += n
	}
	return
}

func nx(k int) int {
	return sort.Search(5*k, func(x int) bool { return zeta(x) >= k })
}
