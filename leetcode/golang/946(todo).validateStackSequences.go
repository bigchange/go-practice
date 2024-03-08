package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	flag := make([]bool, len(pushed))
	i,j := 0,0
	for j <= len(popped) - 1 {
		if pushed[i] == popped[j] {
			flag[i] = true
			j++
			k := i - 1
			for k > 0 {

			}
		} else {
			i++
		}
	}
	return  true
}

func nextCheck(pushed []int, popped []int, flag []bool, i,j int) bool {
  return false
}
