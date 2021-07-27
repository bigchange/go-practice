package main

func IsValid(s string) bool {
	getLetter := false
	getNum := false
	for i, c := range s {
		if i == 0 {
			if !IsLetter(c) {
				return false
			}
			getLetter = true
		}

		if !IsLetter(c) && !IsNumeric(c) {
			return false
		}
		if IsNumeric(c) {
			getNum = true
		}
	}
	if getLetter && getNum {
		return true
	}
	return false
}

func IsLetter(c rune) bool {
	if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
		return true
	}
	return false
}
func IsNumeric(c rune) bool {
	if '0' <= c && c <= '9' {
		return true
	}
	return false
}
