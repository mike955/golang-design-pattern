package main

import "fmt"

// 桥接模式: 将抽象部分与实现部分分离，使它们可以独立变化

/*
	有两台电脑 mac 和 windows，两台打印机 espon 和 hp，两台电脑和打印机可能会任意组合，如果引入
	新的打印机，不希望代码成倍增长，创建两个层次结构:
		抽象层: 计算机
		实现层: 打印机
*/

type computer interface {
	print()
	setPrinter(printer)
}

type mac struct {
	printer printer
}

func (m *mac) print(ctx string) string {
	return "mac_" + m.printer.printContext(ctx)
}

type Windows struct {
	printer printer
}

func (w *Windows) print(ctx string) string {
	return "windows_" + w.printer.printContext(ctx)
}

type printer interface {
	printContext(ctx string) string
}

type espon struct {
}

func (e *espon) printContext(ctx string) string {
	return "espon_" + ctx
}

type hp struct {
}

func (h *hp) printContext(ctx string) string {
	return "hp_" + ctx
}

func main() {
	espon := &espon{}
	hp := &hp{}
	m := &mac{
		printer: espon,
	}

	w := &Windows{
		printer: hp,
	}

	fmt.Println(m.print("hello")) // mac_espon_hello
	fmt.Println(w.print("hello")) // windows_hp_hello
}
