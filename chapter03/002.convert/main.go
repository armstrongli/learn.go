package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	convertType()
	fmt.Println("finish")
}

func convertType() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic啦:", r)
			debug.PrintStack()
		}
	}()
	var a interface{} // any
	a = "string aaa"

	b := a.(int)
	fmt.Println(b)
}
