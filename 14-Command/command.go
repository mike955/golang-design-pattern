package main

import (
	"fmt"
	"strings"
)

/* 命令模式: 将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化，对请求排队或者记录请求日志，以及支持可取消的操作
 *	- 将请求封装为独立对象，创建的对象拥有有关请求的所有信息，可以单独被执行
 *	- 可以将命令作为方法的参数进行传递、将命令保存在其他对象中，或者在运行时切换已连接的命令等
 */

type Commander interface {
	execute() string
}

type command1 struct {
	message string
}

func (c command1) execute() string {
	return c.message
}

func NewCommand(message string) *command1 {
	return &command1{
		message: message,
	}
}

type command2 struct {
	message string
}

func (c command2) execute() string {
	return c.message
}

type CommandQueuer struct {
	queue []Commander
}

func (p *CommandQueuer) addCommand(c Commander) {
	p.queue = append(p.queue, c)
}

func (p *CommandQueuer) execute() string {
	var res []string
	for _, command := range p.queue {
		res = append(res, command.execute())
	}
	return strings.Join(res, " | ")
}

func main() {
	queue := &CommandQueuer{}
	queue.addCommand(NewCommand("first command"))
	queue.addCommand(NewCommand("second command"))
	queue.addCommand(NewCommand("first command"))
	fmt.Println(queue.execute()) // first command | second command | first command
}
