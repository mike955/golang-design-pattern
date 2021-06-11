package main

import "fmt"

/* 原型模式: 用原型实例指定创建对象的种类，并且通过拷贝这个原型来创建新的对象
 * 	- 原型模式能够复制对象，而无需使代码依赖它们所属的类
 *	- 所有原型类型都必须有一个通用的接口，使得即使在对象所属的具体类未知的情况下也能复制对象
 *	- 原型对象可以生成自身的完整副本
 */

type Human interface {
	name() string
	clone() Human
}

type man struct {
	msg string
}

func (f *man) name() string {
	return f.msg
}

func (f *man) clone() Human {
	return &man{
		msg: f.msg + "_clone",
	}
}

func main() {
	mike := &man{msg: "mike"}
	tom := mike.clone()

	fmt.Println(mike.msg)   // mike
	fmt.Println(tom.name()) // mike_clone
}
