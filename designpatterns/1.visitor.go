/*
 * @Author: Jerry You
 * @CreatedDate: 2023/4/17
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/4/17 15:41
 */

// 访问者模式: 一个或多个操作应用到一组对象，操作和对象解耦

package designpatterns

import "fmt"

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

// VisitorFunc 访问者的操作函数和对象的定义
type VisitorFunc func(*Info, error) error

// Visitor 访问者接口定义
type Visitor interface {
	Visit(VisitorFunc) error
}

// Visit 对象实现访问者接口
func (info *Info) Visit(fn VisitorFunc) error {
	// 每个VisitorFunc参数中传入的info，就是这里传入的并执行
	// 根据 v = OtherV(NameV(LogV(Info自身)))， v.Visit(f)，调用Info的Visitor为LogVisitor，此处fn的声明在LogVisitor中
	// 传入参数（info， nil），真的fn执行，然后开始函数嵌套调用
	return fn(info, nil)
}

// DecoratedVisitor 访问者模式下的装饰器
// 和一般模式下访问者封装使用调用的顺序存在差异， 参考测试用例：TestNormalVisitor 和 TestDecoratedVisitor_Visit
type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

// NewDecoratedVisitor 构造函数
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
		//  最先执行最外层传入的f
		if err := fn(info, nil); err != nil {
			return err
		}
		// 然后在装饰器写入的顺序，调用VisitorFunc并执行
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

// Visitor 嵌套使用

// NameVisitor name
type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("NameVisitor() before call function")
		// 根据 v = OtherV(NameV(LogV(Info自身)))， v.Visit(f)， 调用NameVisitor的Visitor为OtherV， 此处fn的声明在OtherV中
		err = fn(info, err)
		// 操作对象字段
		// todo ...
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

type OtherThingsVisitor struct {
	visitor Visitor
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("OtherThingsVisitor() before call function")
		// 根据 v = OtherV(NameV(LogV(Info自身)))， v.Visit(f)， 调用OtherThingsVisitor的Visitor为v， 此处fn的声明为f
		// 执行对应的f函数， 然后递归返回到最开始调用的地方
		err = fn(info, err)
		// 操作对象字段
		// todo ...
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

type LogVisitor struct {
	visitor Visitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("LogVisitor() before call function")
		// 根据 v = OtherV(NameV(LogV(Info自身)))， v.Visit(f)，调用LogVisitor的Visitor为NameVisitor，此处fn的声明在NameVisitor中
		err = fn(info, err)
		// log
		fmt.Printf("==> Name=%s, NameSpace=%s, other=%s\n", info.Name, info.Namespace, info.OtherThings)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}
