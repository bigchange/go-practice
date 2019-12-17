package main

import "fmt"

/*
 * @Author: lolaliva
 * @CreatedDate: 2019-12-17 09:53
 * @Last Modified by: lolaliva
 * @Last Modified time: 2019-12-17 09:53
 */

// 1. struct 嵌入
type Car struct {}

func (c * Car) UseA() {
	fmt.Println("car useA")
	c.UseB()
}

func (c *Car) UseB() {
	fmt.Println("car useB")
}

type BMW struct {
	// struct的嵌入，代表内部声明了一个其他struct的变量
	Car
}

func (b *BMW) UseB() {
	fmt.Println("bmw useB")
}

func main() {
	var bmw = BMW{}
	bmw.UseA() // bmw.Car.UseA()
}
