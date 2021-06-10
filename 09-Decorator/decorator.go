package main

import "fmt"

// 装饰器模式: 动态的给一个对象添加一些额外的职责，就扩展功能而言，装饰器比生成子类方法更为灵活
// 装饰子类的方法，添加一个额外的功能

type human interface {
	Age() int
}

type child struct {
	age int
}

func (c *child) Age() int {
	return c.age
}

type adult struct {
	child human
}

func (a *adult) Age() int {
	age := a.child.Age() + 10 // 对 age 方法进行装饰
	return age
}

func main() {
	mike := &child{age: 10}
	tom := &adult{child: mike}

	fmt.Println(mike.Age()) // 10
	fmt.Println(tom.Age())  // 20
}
