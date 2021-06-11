package main

import "fmt"

// 外观模式: 为子系统中的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用
// 	- 隐藏底层系统的复杂性，为系统中底层的更多个接口提供一个统一的接口

type man struct {
	age    int
	gender string
}

func NewMan(age int) *man {
	return &man{
		age:    age,
		gender: "male",
	}
}

type women struct {
	age    int
	gender string
}

func NewWomen(age int) *women {
	return &women{
		age:    age,
		gender: "female",
	}
}

type human struct {
	man   *man
	women *women
}

func NewHuman(male, female int) *human {
	return &human{
		man: &man{
			age:    male,
			gender: "male",
		},
		women: &women{
			age:    female,
			gender: "female",
		},
	}
}

func (h *human) manAge() int {
	return h.man.age
}

func (h *human) womenAge() int {
	return h.women.age
}

func main() {
	h := NewHuman(20, 18)
	fmt.Println(h.manAge())   // 20
	fmt.Println(h.womenAge()) // 18
}
