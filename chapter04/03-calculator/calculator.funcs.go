package main

import (
	"fmt"
)

func (c Calculator) Add() int {
	fmt.Printf("&c @ calculator.funcs =%p\n", &c)
	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Println("调用Add函数，c的result=", c.result)
	return tempResult
}
func (c Calculator) Sub() int {
	return 0
}
func (c Calculator) Multiple() int {
	return 0
}
func (c Calculator) Divide() int {
	return 0
}
func (c Calculator) Reminder() int {
	return c.left % c.right
}

func (c *Calculator) SetResult(result int) {
	c.result = result
}
