package main

import "fmt"

// 适配器模式: 将一个类的接口转换为客户希望的另一个接口，适配器模式将原本由于接口不兼容而不能一起工作的类可以一起工作

// 适配器模式作为两个不兼容接口之间的桥梁，通过封装对象将复杂的转化过程隐藏在幕后，适配器中持有旧接口对象，并实现新接口

type OldPlayer interface {
	playMp3(fileName string) string
	playWma(fileName string) string
}

type oldPlayer struct {
}

func (o *oldPlayer) playMp3(fileName string) string {
	return "mp3_" + fileName
}

func (o *oldPlayer) playWma(fileName string) string {
	return "wma_" + fileName
}

type NewPlayer interface {
	playMp4(fileName string) string
}

type newPlayer struct {
	oldPlayer OldPlayer
}

func (n *newPlayer) playMp4(fileName string) string {
	return "mp4_" + fileName
}

func main() {
	oldPlayer := &oldPlayer{}
	newPlayer := newPlayer{
		oldPlayer,
	}
	fmt.Println(newPlayer.oldPlayer.playMp3("Singing to remember"))
	fmt.Println(newPlayer.oldPlayer.playWma("Singing to remember"))
	fmt.Println(newPlayer.playMp4("Singing to remember"))
}
