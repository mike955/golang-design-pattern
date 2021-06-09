package main

// 建造者: 将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示

type Builder interface {
	Age() int
	Gender() string
}

type human struct {
	builder Builder
}

func newHuman(b Builder) *human {
	return &human{
		builder: b
	}
}

func (f *human) SetBuilder(b Builder) {
	f.builder = b
}


type manBuilder  struct {
	age int
}

func (m *manBuilder) Age() int {
	return man.age
}

type womenBuilder  struct {
	age int
}

func (w *womenBuilder) Age() int {
	return women.age
}