package main

import (
	"fmt"
)

func panicAndRecover() {
	defer coverPanicUpgraded() // 这样是能抓住严重错误的

	defer func() { // 抓不住严重错误
		coverPanicUpgraded()
	}()

	defer coverPanic()
	var nameScore map[string]int = nil

	nameScore["进文"] = 100
}

func coverPanic() { // 未能抓住panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出严重故障：", r)
		}
	}()
}

func coverPanicUpgraded() {
	if r := recover(); r != nil {
		fmt.Println("系统出严重故障：", r)
	}
}
