package main

import "fmt"

/* 工厂方法: 定义一个用于创建对象的接口，让子类决定将哪一个类实例化，使一个类的实例化延迟到子类
 *	- 客户端只与工厂结构交互并告知需要创建的实例类型，工厂类与相应的具体结构交互，并返回正确的实例
 *	- 该模式提供一个创建对象的方法，允许子类决定实例化对象的类型
 */

type human interface {
	Age() int
	Height() int
	Gender() string
}

type man struct { // 具体产品
	age    int
	height int
	gender string
}

func (m *man) Age() int {
	return m.age
}

func (m *man) Height() int {
	return m.height
}

func (m *man) Gender() string {
	return m.gender
}

type women struct { // 具体产品
	age    int
	height int
	gender string
}

func (w *women) Age() int {
	return w.age
}

func (w *women) Height() int {
	return w.height
}

func (w *women) Gender() string {
	return w.gender
}

func getHuman(gender string) human {
	switch gender {
	case "man":
		return &man{age: 20}
	case "women":
		return &women{age: 18}
	default:
		return nil
	}
}

func main() {
	man := getHuman("man")
	women := getHuman("women")

	fmt.Println(man.Age())   // 20
	fmt.Println(women.Age()) // 18
}
