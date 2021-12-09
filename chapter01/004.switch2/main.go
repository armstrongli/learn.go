package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money interface{} = 10
	// var money interface{} = 10.0
	// var money interface{} = "10"
	switch newMoney := money.(type) {
	case int:
		tmpMoney := newMoney + 3.0
		fmt.Println("money是 int", tmpMoney)
	case int64:
		tmpMoney := newMoney + 3.0
		fmt.Println("money是 int64", tmpMoney)
	case float64:
		tmpMoney := newMoney + 3.1234
		fmt.Println("money是 float64", tmpMoney)
	case float32:
		tmpMoney := newMoney + 3.1234
		fmt.Println("money是 float32", tmpMoney)
	default:
		fmt.Println("money是未知类型")
	}
	fmt.Println("结束", money, reflect.TypeOf(money))
	money = 3
	fmt.Println("结束", money, reflect.TypeOf(money))
}
