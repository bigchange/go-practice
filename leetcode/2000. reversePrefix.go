package leetcode

func reversePrefix(word string, ch byte) string {
	last := -1
	for i, w := range word {
		if byte(w) == ch {
			last = i
			break
		}
	}
	if last == -1 {
		return word
	}
	var s []byte
	for  i := last; i >= 0; i-- {
		s = append(s, word[i])
	}
	return string(s) + word[last+1:]
}
