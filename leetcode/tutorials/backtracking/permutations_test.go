package backtracking

import (
	"reflect"
	"testing"
)

func Test_permutationsII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"2 items", args{nums: []int{1, 2}}, [][]int{{1, 2}, {2, 1}}},
		{"3 items", args{nums: []int{1, 2, 3}}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutationsII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutationsII() = %v, want %v", got, tt.want)
			}
			if got := permutationsI(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutationsI() = %v, want %v", got, tt.want)
			}
		})
	}
}
