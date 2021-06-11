package main

import "fmt"

/* 策略模式: 定义一系列的算法，把它们一个个封装起来，并且使它们可相互替换，该模式将算法的变化独立于使用它的客户端
 *	- 策略模式将一组行为转换为对象，使其在原始上下文(对象)内部能够相互替换
 *	- 原始对象被称为上下文，它包含指向策略对象的引用并将执行行为的任务分派给策略对象
 *	- 为了改变上下文完成其工作方式，其他对象可以使用另一个对象来替换当前链接的策略对象
 *	- 策略模式建议从黄建一个算法类，每个算法都有自己的类，类中的每一个都遵循相同的接口，使得算法可以互相替换
 */

type Algo interface {
	evict(c *cache)
}

type fifo struct{}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type lru struct{}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

type lfu struct{}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

type cache struct {
	storage  map[string]string
	algo     Algo
	capacity int
}

func initCache(e Algo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:  storage,
		capacity: 10,
	}
}

func (c *cache) add(key, value string) {
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) setAlgo(e Algo) {
	c.algo = e
}

func main() {
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.add("a", "1")
}
