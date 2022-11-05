/*
 * @Author: Jerry You
 * @CreatedDate: 2022/11/5
 * @Last Modified by: lolaliva
 * @Last Modified time: 2022/11/5 19:43
 */
package leetcode

import "testing"

func Test_parseBoolExpr(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"!(f)"}, true},
		{"2", args{"|(f,t)"}, true},
		{"3", args{"&(t,f)"}, false},
		{"4", args{"|(&(t,f,t),!(t))"}, false},
		{"5", args{"!(&(f,t))"}, true},
		{"6", args{"&(|(f))"}, false},
		{"7",args{"&(t,t,t)"}, true},
		{"8", args{"|(f,&(t,t))"}, true},
		{"9", args{"!(&(!(&(f)),&(t),|(f,f,t)))"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBoolExpr(tt.args.expression); got != tt.want {
				t.Errorf("parseBoolExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
