package main

import "fmt"

/* 状态模式: 允许一个对象在其内部状态改变时改变它的行为，对象看起来似乎修改了它所属的类
 *	- 在一个对象的内部状态改变时改变其行为
 */

type state interface {
	add(int) int
	consume(int) int
}

type human struct {
	money  int
	gender string

	currentState state
}

func newHuman(money int, gender string) *human {
	h := &human{
		money:  money,
		gender: gender,
	}
	h.currentState = &man{}
	return h
}

func (h *human) add(n int) int {
	h.money = h.money + h.currentState.add(n)
	return h.money
}

func (h *human) consume(n int) int {
	h.money = h.money - h.currentState.consume(n)
	return h.money
}

func (h *human) setState(s state) {
	h.currentState = s
}

type man struct {
}

func (m *man) add(n int) int {
	return n + 10
}

func (m *man) consume(n int) int {
	return n + 10
}

type women struct {
	human *human
}

func (w *women) add(n int) int {
	return n
}

func (w *women) consume(n int) int {
	return n
}

func main() {
	human := newHuman(0, "male")
	fmt.Println(human.money) // 0

	human.add(10)
	fmt.Println(human.money) // 20

	human.consume(10)
	fmt.Println(human.money) // 0

	human.setState(&women{})
	human.add(10)
	fmt.Println(human.money) // 10

	human.consume(10)
	fmt.Println(human.money) // 0
}
