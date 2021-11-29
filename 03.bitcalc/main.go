package main

import (
	"fmt"
)

func main() {
	a, b := 100, 31
	fmt.Println(a ^ b)
	fmt.Println(b ^ a)

	arr := []int{4, 3, 4, 5, 6, 7, 3, 5, 6}
	result := -1
	for _, item := range arr {
		if result < 0 {
			result = item
		} else {
			result = result ^ item
		}
	}
	fmt.Println(result)
}
