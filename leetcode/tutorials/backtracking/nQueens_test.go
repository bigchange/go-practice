package backtracking

import (
	"reflect"
	"testing"
)

func Test_nQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][][]string
	}{
		// TODO: Add test cases.
		{"n=3", args{3}, [][][]string{}},
		{"n=4", args{4}, [][][]string{
			{
				{"#", "Q", "#", "#"},
				{"#", "#", "#", "Q"},
				{"Q", "#", "#", "#"},
				{"#", "#", "Q", "#"},
			},
			{
				{"#", "#", "Q", "#"},
				{"Q", "#", "#", "#"},
				{"#", "#", "#", "Q"},
				{"#", "Q", "#", "#"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
