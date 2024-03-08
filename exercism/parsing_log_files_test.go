/*
 * @Author: Jerry You
 * @CreatedDate: 2023/3/15
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/3/15 13:04
 */
package exercism

import (
	"reflect"
	"testing"
)

func TestIsValidLine(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "invalid ERR message", args: args{"Any old [ERR] text"}, want: false},
		{name: "valid error message", args: args{"[ERR] a good error message"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidLine(tt.args.text); got != tt.want {
				t.Errorf("IsValidLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitLogLine(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{name: "three sections",
			args: args{"section 1<*>section 2<~~~>section 3"}, want: []string{"section 1", "section 2", "section 3"}},
		{name: "three sections with different symbols inside angular brackets",
			args: args{"section 1<=>section 2<-*~*->section 3"}, want: []string{"section 1", "section 2", "section 3"}},
		{name: "two sections with no characters between angular brackets", args: args{"section 1<>section 2"}, want: []string{"section 1", "section 2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitLogLine(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitLogLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountQuotedPasswords(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name:"text with two matches", args: args{lines: []string{``,
			`[INF] passWord`,
			`"passWord"`,
			`[INF] User saw error message "Unexpected Error" on page load.`,
			`[INF] The message "Please reset your password" was ignored by the user`,
		}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountQuotedPasswords(tt.args.lines); got != tt.want {
				t.Errorf("CountQuotedPasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveEndOfLineText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{ name: "INF message",
			args: args{       "[INF] end-of-line23033 Network Failure end-of-line27"},
			want:    "[INF]  Network Failure ",},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveEndOfLineText(tt.args.text); got != tt.want {
				t.Errorf("RemoveEndOfLineText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagWithUserName(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{name: "withUser", args: args{[]string{
			"[WRN] User James123 has exceeded storage space.",
			"[WRN] Host down. User   Michelle4 lost connection.",

		}}, want: []string{
			"[USR] James123 [WRN] User James123 has exceeded storage space.",
			"[USR] Michelle4 [WRN] Host down. User   Michelle4 lost connection.",
		}},
		{name: "without user", args: args{[]string{
			"[INF] Users can login again after 23:00.",
			"[DBG] We need to check that user names are at least 6 chars long.",
		}}, want: []string{
			"[INF] Users can login again after 23:00.",
			"[DBG] We need to check that user names are at least 6 chars long.",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TagWithUserName(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagWithUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}