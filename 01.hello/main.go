package main

import (
	"fmt"
)

func main() {
	var name string = "小强"
	fmt.Println(name)
	var age int = 35
	fmt.Println(age)
	var tall float64 = 1.70
	fmt.Println(tall)

	name = "小强测试变量"
	// age = "测试一下类型" // 类型不兼容报错
	// tall = age // 类型不兼容错误
}
