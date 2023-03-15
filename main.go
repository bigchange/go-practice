/*
 * @Author: lolaliva
 * @CreatedDate: 2019-12-17 09:53
 * @Last Modified by: lolaliva
 * @Last Modified time: 2019-12-17 09:53
 */

package main

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bigchange/go-practice/leetcode/golang"
)

func tutorialsEntry() {
}

func leetCodeEntry() {
	leetcode.Entry()
}

func main() {
	leetCodeEntry()
	TestF()
	// meituan()
	// TestSyncPool()
	// DecoratedVisitorLoadFile()
	// NormalVisitor()
	tutorialsEntry()

}

func meituan() {
	var n int
	fmt.Scanln(&n)
	var strs []string
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanln(&s)
		strs = append(strs, s)
	}
	for i := 0; i < n; i++ {
		if IsValid(strs[i]) {
			fmt.Println("Accept")
			continue
		}
		fmt.Println("Wrong")
	}
}

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

// -------- testF start ---- //

type V struct {
	A string
	B string
}

func (v V) setA(a string) {
	v.A = a
}
func (v V) setB(b string) {
	v.B = b
}

func getP() {
	fmt.Println("NumCPU:", runtime.GOMAXPROCS(runtime.NumCPU()))
}

var value int32

func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, v + delta) {
			fmt.Println("value:", value)
			break
		}
	}
}
type MyMutex struct {
	count int
	sync.Mutex
}
func Mutex() {
	var mu MyMutex
	mu.Lock()
	var mu2 = mu
	mu.count++
	mu.Unlock()
	fmt.Println(mu.count, mu2.count)
	mu2.Lock()
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}
func Block() {
	var i byte
	runtime.GOMAXPROCS(1)
	go func() {
		for i = 0; i <= 255; i++ {

		}
	}()
	runtime.Gosched()
	runtime.GC()

  // panic
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()

}

var poolNew = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}
func PoolTest() {
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}
func processRequest(size int) {
	b := poolNew.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	poolNew.Put(b)
	time.Sleep(1 * time.Millisecond)
}

func TestF() {
	//
	v := V{A: "A", B: "B"}
	v.setA("a")
	v.setB("b")
	fmt.Printf("A:%v, B:%v\n", v.A, v.B)

	getP()
	SetValue(100)
	// Block()
	// Mutex()
	// PoolTest()
	var one = new(Once)
	var wg sync.WaitGroup
	runNum := 1000
	wg.Add(runNum)
	for i:= 0; i < runNum; i++ {
		go func() {
			one.Do(func() {
				fmt.Println("init once fun logic")
			})
			wg.Done()
		}()
	}
	wg.Wait()
}
type Once struct {
	m    sync.Mutex
	done uint32
}
func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

// -------- testF end ---- //

// GenerateNatural 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				fmt.Printf("generated %v\n", i)
			}
		}
	}()
	return ch
}

// PrimeFilter 管道过滤器: 删除能被素数整除的数
func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
					fmt.Printf("output %v\n", i)
				}
			}
		}
	}()
	return out
}

var pool *sync.Pool

type Person struct {
	Name string
	noCopy
}

type noCopy struct{}

// Lock 是一个空操作用来给 `go vet` 的 -copylocks 静态分析
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// Shape 接口完整性检查: 强验证
type Shape interface {
	String()
}

type Circle struct{}

type Rectangle struct {
	Square
	// s Square // compile error
}

type Square struct {
}

func (s Square) String() {
	fmt.Println("this is a square")
}

// compile error
// var _ Shape = (*Circle)(nil)
// compile passed
var _ Shape = (*Square)(nil)
var _ Shape = (*Rectangle)(nil)

func noCopyPerson(p *Person) {
	p.Name = "noCopyPerson"
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
	runtime.GOMAXPROCS(2)
	p := pool.Get().(*Person)
	fmt.Printf("首次从 pool 里获取：%v \n", p)
	// 使用 go vet . 命令可以检测出来错误（有输出的就是错误）
	// copy := func(p Person) Person {
	// 	p.Name = "copied"
	// 	return p
	// }
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

	for i := 0; i < 4; i++ {
		fmt.Printf("Pool 调用 Get: %v \n", pool.Get().(*Person))
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
	v = NewDecoratedVisitor(v, NameVisitorFunc, LogVisitorFunc, OtherThingsVisitorFunc)
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
