package leetcode

func lengthOfLastWord(s string) int {
	l := len(s)
	c := 0
	for i := l - 1; i >=0; i-- {
		if c > 0 {
			if s[i] == ' ' {
				return c
			}
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			c++
		}
		if  s[i] >= 'a' && s[i] <= 'z' {
			c++
		}
		continue
	}
	return c
}