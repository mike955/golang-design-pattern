package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* 模版方法模式: 定义一个操作中的算法的骨架，而将一些步骤延迟到子类中，该模式使得子类可以不改变一个算法的结构即可重定义该算法的某些特定操步骤
 *	- 该模式在基类中定义了一个算法框架，允许子类在不修改结构的情况下重写算法的特定步骤
 *	- 如果某个特定操作的步骤是相同的，但实现方式有所不同，则适合使用该模式
 */

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type genPassword interface {
	genPassword(int) string
	gensalt(int) string
	encrypt(string, string) string
}

type gen struct {
	gen genPassword
}

func (u *gen) encrypt(p, s string) string {
	return u.gen.encrypt(p, s)
}

type user1 struct {
}

func (u *user1) genPassword(d int) string {
	return fmt.Sprintf("%d", rand.Intn(100))
}

func (u *user1) gensalt(d int) string {
	return fmt.Sprintf("%d", rand.Intn(100))
}

func (u *user1) encrypt(p, s string) string {
	return p + s
}

type user2 struct {
}

func (u *user2) genPassword(d int) string {
	b := make([]byte, d)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (u *user2) gensalt(d int) string {
	b := make([]byte, d)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (u *user2) encrypt(p, s string) string {
	return p + s
}

func main() {
	rand.Seed(time.Now().UnixNano())

	user1 := &user1{}
	user2 := &user2{}

	gen := &gen{gen: user1}
	fmt.Println(gen.encrypt(gen.gen.genPassword(2), gen.gen.gensalt(2))) //	9392
	gen.gen = user2
	fmt.Println(gen.encrypt(gen.gen.genPassword(2), gen.gen.gensalt(2))) //	GdzS
}
