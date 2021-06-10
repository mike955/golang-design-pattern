package main

import (
	"errors"
	"fmt"
)

/* 代理模式: 为其它对象提供给一个代理以控制这个对象的访问
- 创建一个新的代理类，实现与主对象相同的接口
*/

type server interface {
	handleRequest(string) (int, error)
}

type application struct{}

func (a *application) handleRequest(url string) (int, error) {
	if url == "/golang/proxy" {
		return 200, nil
	}
	return 404, errors.New("NOT FOUND")
}

type proxy struct {
	application server
	rateLimiter map[string]int
}

func NewProxy(app server) *proxy {
	return &proxy{
		application: app,
		rateLimiter: make(map[string]int),
	}
}

func (p *proxy) handleRequest(url string) (int, error) {
	if p.rateLimiter[url] > 10 {
		return 403, errors.New("limit error")
	}
	p.rateLimiter[url] = p.rateLimiter[url] + 1
	return p.application.handleRequest(url)
}

func main() {
	app := &application{}
	proxy := NewProxy(app)
	code, err := proxy.handleRequest("/golang/proxy")
	fmt.Println(code) // 200
	fmt.Println(err)  // nil
}
