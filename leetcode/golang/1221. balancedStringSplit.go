package leetcode


func balancedStringSplit(s string) int {
	LC, RC := 0,0
	c := 0
	for _, v := range s {
		if v == 'L' {
			LC++
		}
		if v == 'R' {
			RC++
		}
		if LC == RC {
			c++
		}
	}
	return c
}
