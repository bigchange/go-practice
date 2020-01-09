/*
 * @Author: lolaliva
 * @CreatedDate: 2019-12-17 09:53
 * @Last Modified by: lolaliva
 * @Last Modified time: 2019-12-17 09:53
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
	noCopy
}

type noCopy struct{}

func main() {
	TestSyncPool()
}

// Lock 是一个空操作用来给 `go vet` 的 -copylocks 静态分析
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return new(Person)
		},
	}
}

func noCopyPerson(p *Person) {
	p.Name = "copyPerson"
}
func copyPerson(p Person) Person {
	p.Name = "copyPerson"
	return p
}

func TestSyncPool() {
	initPool()
	runtime.GOMAXPROCS(0)
	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)
	copyPerson(*p)
	noCopyPerson(p)
	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	pool.Put(p)

	p1 := &Person{Name: "second"}
	pool.Put(p1)

	p2 := &Person{Name: "third"}
	pool.Put(p2)

	for i := 0; i < 10; i++ {
		fmt.Println("Pool 调用 Get: ", pool.Get().(*Person))
	}
}
