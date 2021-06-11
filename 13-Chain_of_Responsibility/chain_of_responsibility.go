package main

import "fmt"

/* 职责链: 为接触请求的发送者和接收者之间的耦合，而使多个对象都有机会处理这个请求，将这些对象链成一条链，并沿着这条链传递该请求，直到有一个对象处理它
 *	- 将请求沿着处理者链传递，每个处理者都可以对请求进行处理
 *	- 类似于中间件模式
 */

type handler interface {
	setNext(handler)
	execute(string) string
}

type firstHanlder struct {
	next handler
}

func (f *firstHanlder) setNext(h handler) {
	f.next = h
}

func (f *firstHanlder) execute(s string) string {
	if f.next != nil {
		return f.next.execute("first handle | " + s)
	}
	return "first handle | " + s
}

type secondHanlder struct {
	next handler
}

func (f *secondHanlder) setNext(h handler) {
	f.next = h
}

func (f *secondHanlder) execute(s string) string {
	if f.next != nil {
		return f.next.execute("second handle | " + s)
	}
	return "second handle | " + s
}

func main() {
	first := &firstHanlder{}
	second := &secondHanlder{}
	second.setNext(first)
	fmt.Println(second.execute("hello")) // first handle | second handle | hello
}
