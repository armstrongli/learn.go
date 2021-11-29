package main

import (
	"fmt"
)

func main() {
	var val interface{} = 10
	switch newVal := val.(type) {
	case int:
		tmpVal := newVal + 3.0
		fmt.Println("int", tmpVal)
	case float64:
		tmpVal := newVal + 3.1234
		fmt.Println("float64", tmpVal)
	default:
		fmt.Println("money是未知类型")
	}
}
