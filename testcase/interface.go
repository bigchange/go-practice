package testcase

import "fmt"

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	// ...
	ret := f.Bar(99)
	fmt.Println(ret)
}


