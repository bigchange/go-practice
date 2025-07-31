package leetcode

import "fmt"

func print(res [][]int) {
	row := res[0]
	rowLen := len(res)
	colLen := len(row)
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			fmt.Printf("%v ", res[i][j])
		}
		fmt.Println()
	}
}
