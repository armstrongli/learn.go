package main

import (
	"fmt"
)

func main() {
	hello()
	helloToSomeone("Dandy")
	helloToSomeone("昊楠")
	msg := constructMessage("昊楠")
	fmt.Println(msg)
}

func hello() {
	fmt.Println("你好，Golang")
}

func helloToSomeone(name string) {
	fmt.Println("你好，", name)
}

func constructMessage(name string) string {
	return "你好，" + name + "，再来一个"
}

func calcBMI(tall float64, weight float64) float64 {
	return tall / (weight * weight)
}
