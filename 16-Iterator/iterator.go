package main

import "fmt"

/* 迭代器模式: 提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露该对象的内部表示
 * - 集合结构提供了一个迭代器，它允许在不暴露其底层实现的前提下，按顺序遍历集合结构中的每个元素
 */

type user struct {
	name string
	age  int
}

type iterator interface {
	hasNext() bool
	getNext() *user
}

type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

func main() {
	user1 := &user{
		name: "mike",
		age:  20,
	}
	user2 := &user{
		name: "tom",
		age:  18,
	}

	userCollection := &userCollection{users: []*user{user1, user2}}
	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		fmt.Println(iterator.getNext().name) // mike \n tom
	}
}
