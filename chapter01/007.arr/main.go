package main

import (
	"fmt"
)

func main() {
	var age int = 35
	fmt.Println(age)
	var ages [5]int = [5]int{35, 32, 33, 37, 59}
	fmt.Println(ages)
	var ages2 = [5]int{35, 32, 33, 37, 59}
	fmt.Println(ages2)
	ages3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(ages3)
	ages2 = ages3
	// ages3 = [6]int{35, 32, 33, 37, 59, 99} // 错误：数组长度不能变

	// var xxxx type
	var ages4 [3]int
	fmt.Println("ages4: ", ages4)
	ages4[0] = 1000
	ages4[1] = 1111
	ages4[2] = 2222
	fmt.Println("ages4: ", ages4)
	// ages4[-1]=-1 //错误，越界
	// ages4[99]=-1 //错误，越界

	for i := 0; i < len(ages4); i++ {
		fmt.Println(ages4[i])
	}

	for i, val := range ages4 {
		fmt.Println(ages4[i], "====>", i, "->", val)
	}
}
