/*
 * @Author: lolaliva
 * @CreatedDate: 2019-12-17 09:53
 * @Last Modified by: lolaliva
 * @Last Modified time: 2019-12-17 09:53
 */

package main

import (
	"fmt"

	leetcode "github.com/bigchange/go-practice/leetcode/golang"
)

func tutorialsEntry() {

}

func leetCodeEntry() {
	leetcode.Entry()
}

func main() {
	leetCodeEntry()
	TestF()
	tutorialsEntry()

}

func nextPowOf2(cap int) int {
	if cap < 2 {
		return 2
	}
	if cap&(cap-1) == 0 {
		return cap
	}
	cap |= cap >> 1
	fmt.Println("cap:", cap)
	cap |= cap >> 2
	fmt.Println("cap:", cap)
	cap |= cap >> 4
	fmt.Println("cap:", cap)
	cap |= cap >> 8
	fmt.Println("cap:", cap)
	cap |= cap >> 16
	fmt.Println("cap:", cap)
	return cap + 1
}

func TestF() {
	le := getMaxStrLen("aaaabbbbcccccdddaaabbccdd", 5)
	fmt.Println("maxLen:", le)
}

// -------- testF start ---- //
type T struct {
	A int
	B string
}

// -------- testF end ---- //
