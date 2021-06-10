package main

import (
	"fmt"
	"sync"
)

// 单例模式: 保证一个类只有一个实例，并提供一个访问它的全局访问点

var once sync.Once
var instance *single

type single struct {
	age int
}

func New(age int) *single {
	once.Do(func() {
		instance = &single{
			age: age,
		}
	})
	return instance
}

func main() {
	var s *single
	for i := 0; i < 10; i++ {
		s = New(i)
	}
	fmt.Println(s.age) // 0
}
