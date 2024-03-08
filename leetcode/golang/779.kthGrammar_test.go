package leetcode

import "testing"

func Test_kthGrammar(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"one",
			args{n: 1, k: 1},
			0,
		},
		{
			"two",
			args{2,1},
			0,
		},
		{
			"three",
			args{3, 2},
			1,
		},
		{
			"three-2",
			args{3, 3},
			1,
		},
		{
			"four",
			args{4, 4},
			0,
		},
		{
			"four-2",
			args{4, 5},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthGrammar(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kthGrammar() = %v, want %v", got, tt.want)
			}
		})
	}
}
