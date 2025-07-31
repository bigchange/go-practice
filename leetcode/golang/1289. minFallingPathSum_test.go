package leetcode

import "testing"

func Test_minFallingPathSum1289(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "bad case",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum1289_2(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSum1289() = %v, want %v", got, tt.want)
			}
		})
	}
}
