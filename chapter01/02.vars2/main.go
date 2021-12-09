package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "小强"
	{
		val, err := strconv.Atoi(name)
		fmt.Println(val, err)
	}
	age := 30
	fmt.Printf("%p\n", &age)
	age, time := 32, "时间"
	fmt.Printf("%p\n", &age)
	fmt.Println(name, age, time)
	{
		age := 3
		fmt.Printf("%p\n", &age)
	}
}
