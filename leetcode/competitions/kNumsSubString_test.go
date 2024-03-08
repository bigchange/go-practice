package competitions

import "testing"

func Test_getMaxSubStrLen(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"fullMatched", args{"aaabbb", 3}, 6},
		{"subMatched-start", args{"aaabb", 3}, 3},
		{"subMatched-end", args{"aaabbbb", 4}, 4},
		{"normalCase", args{"aaabbccaabbaac", 4}, 6},
		{"notMatched", args{"aaabbb", 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxSubStrLen(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("getMaxSubStrLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
