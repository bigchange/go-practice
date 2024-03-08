/*
 * @Author: Jerry You
 * @CreatedDate: 2023/3/28
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/3/28 11:09
 */
package leetcode

import "testing"

func TestDeterminant(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {name:"3*3", args: args{[][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}}}, want: -20},
		{name: "4*4", args: args{matrix: [][]int{{2, 5, 3, 1}, {1, -2, -1,1}, {1, 3, 4,1},{1, 3, 4,1}}}, want: -100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Determinant(tt.args.matrix); got != tt.want {
				t.Errorf("Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}
