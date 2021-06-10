package main

import "fmt"

// 组成模式: 将对象组成树形接口以表示 "整体-部分" 的层次结构，组成模式是的客户对于单个对象和复合对象的使用具有一致性

type man interface {
	name() string
}

type human struct {
	mans []man
	name string
}

func (h *human) search(name string) bool {
	for _, man := range h.mans {
		if man.name() == name {
			return true
		} else {
			return false
		}
	}
	return false
}

type mike struct {
	n string
}

func (m *mike) name() string {
	return m.n
}

func main() {
	mike := &mike{n: "mike"}
	human := &human{
		name: "human1",
		mans: []man{},
	}
	human.mans = append(human.mans, mike)
	fmt.Println(human.search("mike")) // true
}
