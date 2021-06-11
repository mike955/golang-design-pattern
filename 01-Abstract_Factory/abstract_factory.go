package main

import "fmt"

/* 抽象工厂: 提供一个创建一系列相关或相互依赖对象的接口，而无需指定他们具体的类

 */
type human interface { // 抽象工厂接口
	genMan() iman
	genWomen() iwomen
}

type iman interface { // 抽象接口
	Age() int
	Gender() string
	Beard() bool
}

type man struct {
	age    int
	gender string
	beard  bool
}

func (m *man) Age() int {
	return m.age
}

func (m *man) Gender() string {
	return m.gender
}

func (m *man) Beard() bool {
	return m.beard
}

type iwomen interface { // 抽象接口
	Age() int
	Gender() string
	Height() int
}

type women struct {
	age    int
	gender string
	heigth int
}

func (w *women) Age() int {
	return w.age
}

func (w *women) Gender() string {
	return w.gender
}

func (w *women) Height() int {
	return w.heigth
}

type american struct { // 具体工厂
}

func (a *american) genMan() iman {
	man := &man{
		age:    20,
		gender: "male",
		beard:  true,
	}
	return man
}

func (a *american) genWomen() iwomen {
	women := &women{
		age:    20,
		gender: "female",
		heigth: 170,
	}
	return women
}

type britain struct { // 具体工厂
}

func (a *britain) genMan() iman {
	man := &man{
		age:    21,
		gender: "male",
		beard:  true,
	}
	return man
}

func (a *britain) genWomen() iwomen {
	women := &women{
		age:    20,
		gender: "female",
		heigth: 171,
	}
	return women
}

func getHuman(gender string) (human, error) {
	switch gender {
	case "american":
		return new(american), nil
	case "britain":
		return new(britain), nil
	default:
		return nil, nil
	}
}

func main() {
	american, _ := getHuman("american")
	britain, _ := getHuman("britain")

	man1 := american.genMan()
	man2 := britain.genMan()

	fmt.Println(man1.Age()) // 20
	fmt.Println(man2.Age()) // 21
}
