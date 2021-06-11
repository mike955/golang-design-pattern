package main

import "fmt"

/* 访问者模式:表示一个作用于某个对象结构中的各个元素的操作，它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作
 *	- 该模式允许在不实际修改结构体的情况下结构体添加行为
 */

type visitor interface {
	visitForMessage1(*message1)
	visitForMessage2(*message2)
}

type visitable interface {
	accept(visitor)
}

type messageVisitor struct{}

func (mv *messageVisitor) visitForMessage1(m *message1) { // 修改 m21行为
	m.msg = fmt.Sprintf("%s %s", m.msg, "visitForMessage1")
}
func (mv *messageVisitor) visitForMessage2(m *message2) { // 修改 m2 行为
	m.msg = fmt.Sprintf("%s %s", m.msg, "visitForMessage2")
}

type message1 struct {
	msg string
}

func (m *message1) accept(v visitor) {
	v.visitForMessage1(m)
}

type message2 struct {
	msg string
}

func (m *message2) accept(v visitor) {
	v.visitForMessage2(m)
}

func main() {
	m1 := &message1{msg: "m1"}
	m2 := &message2{msg: "m2"}

	messageVisitor := &messageVisitor{}
	m1.accept(messageVisitor) // 添加访问者修改自身行为
	m2.accept(messageVisitor) // 添加访问者修改自身行为

	fmt.Println(m1.msg) //	m1 visitForMessage1
	fmt.Println(m2.msg) //  m2 visitForMessage2
}
