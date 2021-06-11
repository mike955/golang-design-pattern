package main

/* 观察者模式: 定义对象的一种一对多的依赖关系，以便当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并自动刷新
 *	- 观察者模式用于对事件进行订阅和消费
 */

type observer interface {
	notify(m string)
}

type subject interface {
	addObserver(o observer)
	deleteObserver(o observer)
	notifyAll(m string)
}

type publisher struct {
	observers []observer
}

func (p *publisher) addObserver(o observer) {
	p.observers = append(p.observers, o)
}

func (p *publisher) deleteObserver(o observer) {
	var index int
	for i, observer := range p.observers {
		if observer == o {
			index = i
			break
		}
	}
	p.observers = append(p.observers[:index], p.observers[index+1:]...)
}

func (p *publisher) notifyAll(m string) {
	for _, observer := range p.observers {
		observer.notify(m)
	}
}

type TestObserver struct {
	id      int
	message string
}

func (p *TestObserver) notify(m string) {
	p.message = p.message + " | m"
}

func main() {
	publisher := &publisher{
		observers: make([]observer, 0),
	}

	observer1 := &TestObserver{
		id:      0,
		message: "observer1",
	}

	observer2 := &TestObserver{
		id:      0,
		message: "observer1",
	}

	publisher.addObserver(observer1)
	publisher.addObserver(observer2)

	publisher.notifyAll("notify message")
}
