package main

// 抽象工厂: 提供一个创建一系列相关或相互依赖对象的接口，而无需指定他们具体的类

type human interface {
	Age() int
}

type man struct {
	age int
}

func (m *man) Age() int {
	man := &man{
		age: 20,
	}
	return man.age
}

type women struct {
	age int
}

func (w *women) Age() int {
	women := &women{
		age: 18,
	}
	return women.age
}

func getHuman(gender string) (human, error) {
	switch gender {
	case "man":
		return new(man), nil
	case "women":
		return new(women), nil
	default:
		return nil, nil
	}
}
