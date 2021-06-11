package main

import "fmt"

/* 中介者模式: 用一个中介者来封装一系列的对象交互，中介者使各对象不需要显示地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互
 *	- 中介者模式创建一个中介对象来防止对象之间的直接通信，避免它们之间的直接依赖关系
 */

type mediator interface {
	sendMsg() string
	receiveMsg(string)
}

type medie struct {
	message []string
}

func (m *medie) sendMsg() string {
	if len(m.message) > 0 {
		msg := m.message[0]
		m.message = m.message[1:]
		return msg
	}
	return ""
}

func (m *medie) receiveMsg(msg string) {
	m.message = append(m.message, msg)
}

type producer struct {
	media mediator
}

func (p *producer) sendMsg(msg string) {
	p.media.receiveMsg(msg)
}

type consumer struct {
	media mediator
}

func (p *consumer) receiveMsg() string {
	return p.media.sendMsg()
}

func main() {
	mediator := &medie{}

	producer := &producer{media: mediator}
	consumer := &consumer{media: mediator}

	producer.sendMsg("123")
	fmt.Println(mediator.message)      //	[123]
	fmt.Println(consumer.receiveMsg()) // 123
}
