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

func main() {
	// TestSyncPool()
	// DecoratedVisitorLoadFile()
	// NormalVisitor()
	fmt.Println("Hello, 世界")
	go p()
	q()
	fmt.Println("Hello, background")
}

var pool *sync.Pool

func p() {
	fmt.Println("this is in p func")
	select {

	}
	fmt.Println("this is in p2 func")
}

func q() {
	fmt.Println("this is in q func")
}

type Person struct {
	Name string
	noCopy
}

type noCopy struct{}
// Lock 是一个空操作用来给 `go vet` 的 -copylocks 静态分析
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}


// 接口完整性检查: 强验证
type Shape interface {
	String()
}

type Circle struct {}

type Rectangle struct {
	Square
	// s Square // compile error
}

type Square struct {
}

func (s *Square) String() {
}

// compile error
// var _ Shape = (*Circle)(nil)
// compile passed
var _ Shape = (*Square)(nil)
var _ Shape = (*Rectangle)(nil)


func noCopyPerson(p *Person) {
	p.Name = "noCopyPerson"
}
func copyPerson(p Person) Person {
	p.Name = "copyPerson"
	return p
}

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return new(Person)
		},
	}
}

func TestSyncPool() {
	initPool()
	runtime.GOMAXPROCS(0)
	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)
	copyPerson(*p)
	fmt.Printf("设置 p.Name = %s\n", p.Name)
	noCopyPerson(p)
	fmt.Printf("设置 p.Name = %s\n", p.Name)
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


type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}
func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}


type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}


type OtherThingsVisitor struct {
	visitor Visitor
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}


type LogVisitor struct {
	visitor Visitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func DecoratedVisitorLoadFile() {

		info := Info{}
		var v Visitor = &info

		NameVisitorFunc := func(info *Info, err error) error {
			fmt.Println("NameVisitor() before call function")
			if err == nil {
				fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
			}
			fmt.Println("NameVisitor() after call function")
			return err
		}
		LogVisitorFunc := func(info *Info, err error) error {
			fmt.Println("LogVisitor() before call function")
			fmt.Println("LogVisitor() after call function")
			return err
		}

		OtherThingsVisitorFunc := func(info *Info, err error) error {
			fmt.Println("OtherThingsVisitor() before call function")
			if err == nil {
				fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
			}
			fmt.Println("OtherThingsVisitor() after call function")
			return err
		}

		loadFile := func(info *Info, err error) error {
			info.Name = "Hao Chen"
			info.Namespace = "MegaEase"
			info.OtherThings = "We are running as remote team."
			return nil
		}
	  v = NewDecoratedVisitor(v, NameVisitorFunc,LogVisitorFunc,OtherThingsVisitorFunc)
		v.Visit(loadFile)
}

func NormalVisitor() {
		info := Info{}
		var v Visitor = &info
		v = LogVisitor{v}
		v = NameVisitor{v}
		v = OtherThingsVisitor{v}

		loadFile := func(info *Info, err error) error {
			info.Name = "Hao Chen"
			info.Namespace = "MegaEase"
			info.OtherThings = "We are running as remote team."
			return nil
		}
		v.Visit(loadFile)
}
