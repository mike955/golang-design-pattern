package main

import (
	"errors"
	"fmt"
)

/* 备忘录模式: 在不破坏封装的前提下，捕获一个对象的内部表示，并在该对象之外保存这个状态，这样以后就可将该对象恢复到保存的状态
 *	- 备忘录不会影响它所处理的对象的内部结构，也不会影响快找中保存的数据
 */

type State struct { // 保存的备忘信息
	Description string
}

type memento struct {
	state State
}

type careTaker struct { // 负责人：用于保存备忘录数据
	mementoList []*memento
}

func (c *careTaker) add(m *memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) get(i int) (*memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return &memento{}, errors.New("index error")
	}
	return c.mementoList[i], nil
}

type originator struct {
	state State
}

func (o *originator) NewMemento() *memento {
	return &memento{
		state: o.state,
	}
}

func (o *originator) setState(s State) {
	o.state = s
}

func main() {
	careTaker := &careTaker{
		mementoList: make([]*memento, 0),
	}

	originator := &originator{state: State{Description: "info"}}
	memento := originator.NewMemento()
	careTaker.add(memento)
	memo, err := careTaker.get(0)
	fmt.Println(memo.state) // {info}
	fmt.Println(err)        // nil
}
