package generic

import (
	"fmt"
	"strconv"
)

// Go 泛型泛型语法为中括号
func Print[T any](t T) {
	fmt.Printf("printing type: %T\n", t)
}

// 类型约束： 对类型和方法的限制

// 方法限制，接口约束
type Stringer interface {
	String() string
}

func Stringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// 类型集合
type Types interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// 限定为 Types
func Sub[T Types](t1, t2 T) T {
	return t1 - t2
}

// 在当前Go版本的实现中，接口值持有实例的指针，将非指针的值传递给一个声明为interface{}类型的形参，会有一个装箱的操作，
// 即在内存中，实例的内容在堆栈上，而接口值则是指向实例位置的指针。
// 但是需要注意的是，在泛型中，泛型类型的值不会被装箱。

// 约束元素

// 其中 int 为基础类型
type Integer interface {
	int
}

// 近似约束元素
// 类型别名
type Phone string
type Email string
type Address string

// AnyString 就可以包括 Phone，Email，Address等类型
// 即： 基础类型为string的类型都可以
type AnyString interface {
	~string
}

func Desensitization[T AnyString](str T) string {
	var newStr string
	// Desensitization logic
	// newStr = desensitizationFunc(str)
	return newStr
}

// 联合约束元素
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func SumOfSignedInteger[T SignedInteger](integers []SignedInteger) SignedInteger {
	sum := 0
	for i := range integers {
		sum += i
	}
	return sum
}

// 约束中的可比类型
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type ComparableHasher interface {
	comparable
	Hash() uintptr
}

// 泛型函数
func Filter[T any](f func(T) bool, src []T) []T {
	var dst []T
	for _, v := range src {
		if f(v) {
			dst = append(dst, v)
		}
	}
	return dst
}

// 指针方法
// 在一些项目的实体定义中，某些实体可能带有一些修改自身属性的方法，这种方法带有写的语义，
// 在Go中，这种方法会声明接受者为接收类型的指针

// 定义一个可设置的 int
type Settable int

// 从字符串中设置 *p 的值
func (p *Settable) Set(s string) {
	// 生产场景代码不应该忽略错误
	i, _ := strconv.Atoi(s)
	*p = Settable(i)
}

type Setter2[B any] interface {
	Set(string)
	*B // non-interface type constraint element
}

func FromStrings2[T any, PT Setter2[T]](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		p := PT(&result[i])
		p.Set(v)
	}
	return result
}

func F() {
	nums := FromStrings2[Settable, *Settable]([]string{"1", "2"})
	// 现在 nums 是 []Settable{1, 2}。
	// 也可以使用类型推断  会简单点
	nums = FromStrings2[Settable]([]string{"1", "2"})

	fmt.Println("nums", nums)
}

// 泛型零值
// 一般使用 var zero T
// 有一个对象，包含一个管道属性，可以调用此对象方法压入或弹出数据
type Queue[T any] struct {
	data chan T
}

// 构建新的队列
func NewQueue[T any](size int) Queue[T] {
	return Queue[T]{
		data: make(chan T, size),
	}
}

// 压入数据
func (q Queue[T]) Push(val T) {
	q.data <- val
}

// 弹出数据 ,如果没有数据会被阻塞
func (q Queue[T]) Pop() T {
	d := <-q.data
	return d
}

// 在该代码中，T可以是任何值，包括可能不为nil的值。
// 我们可以利用var语句来解决这个问题，它生成一个新变量，并将其初始化为该类型的零值：
func Zero[T any]() T {
	var zero T
	return zero
}

// 根据这一特性，可以改写TryPop方法
func (q Queue[T]) TryPop() (T, bool) {
	select {
	case val := <-q.data:
		return val, true
	default:
		// 可编译通过
		var zero T
		return zero, false
	}
}
