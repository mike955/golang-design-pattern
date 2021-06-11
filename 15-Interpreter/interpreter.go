package main

import (
	"fmt"
	"strings"
)

/* 解释器模式:给定一个语言，定义它的文法的一种表示，并定义一个解释器，该解释器使用该表示来解释语言中的句子
 */

func interpreter(input string) string {
	switch input {
	case "00":
		return ""
	case "01":
		return "h"
	case "02":
		return "e"
	case "03":
		return "l"
	case "04":
		return "o"
	default:
		return "0"
	}
}

func calculate(input string) string {
	res := ""
	inputs := strings.Split(input, " ")
	for _, in := range inputs {
		res = res + interpreter(in)
	}
	return res
}

func main() {
	str := "01 02 03 03 04"
	res := calculate(str)
	fmt.Println(res) // hello
}
