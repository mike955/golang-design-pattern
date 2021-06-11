package main

import (
	"fmt"
	"net/http"
)

/*
	享元模式: 运用共享技术有效地支持大量细粒度的对象
		- 享元模式会从对象中提取公共部分创建享元对象，这些享元对象可以在多个对象中分享，减少了重复创建
		- 享元模式中，将享元对象存储在 map 容器中，当创建享元对象时从 map 容器中获取享元对象
*/

type httpClient struct {
	url    string
	client *http.Client
}

type httpClientMap struct {
	clientMap map[string]*httpClient
}

func NewHttpClientMap() *httpClientMap {
	return &httpClientMap{
		clientMap: map[string]*httpClient{},
	}
}
func (h *httpClientMap) get(url string) *httpClient {
	if _, ok := h.clientMap[url]; ok {
		return h.clientMap[url]
	}
	client := &httpClient{
		url:    url,
		client: &http.Client{},
	}
	h.clientMap[url] = client
	return client
}

func main() {
	c := NewHttpClientMap()
	client := c.get("https://github.com/mike955")
	fmt.Println(client.url) // https://github.com/mike955
}
