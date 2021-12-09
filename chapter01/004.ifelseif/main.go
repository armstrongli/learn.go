package main

import (
	"fmt"
)

func main() {
	var rmb int = 21
	if rmb <= 20 {
		// 如果有不超过20，点个外卖
		fmt.Println("点个外卖") // a
	} else if rmb <= 200 {
		// 如果不超过200，下个馆子
		fmt.Println("下个馆子") // b
	} else if rmb <= 2000 {
		// 如果不超过2000，去米其林
		fmt.Println("去米其林") // c
	} else if rmb <= 20000 {
		// 如果不超过20000，去新马泰
		fmt.Println("去新马泰") // d
	} else {
		// 如果再多，容我想想
		fmt.Println("容我想想") // e
	}
	fmt.Println("end")
}
