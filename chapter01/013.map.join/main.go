package main

import (
	"fmt"
)

func main() {
	leftMap, rightMap := map[string]int{}, map[string]int{}

	leftMap["语文"] = 80
	rightMap["数学"] = 43

	for k, v := range rightMap {
		leftMap[k] = v
	}
	fmt.Println(leftMap)
}
