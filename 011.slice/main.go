package main

import (
	"fmt"
)

func main() {
	{
		a := make([]int, 0)
		fmt.Println("len: ", len(a))
		a = append(a, 1)
		fmt.Println("len: ", len(a), "val:", a)
	}

	{
		a := make([]int, 1)
		fmt.Println("len: ", len(a), "cap: ", cap(a))
		a = append(a, 1)
		fmt.Println("len: ", len(a), "val:", a)
	}

	{
		a := make([]int, 0, 1)
		fmt.Println("len: ", len(a), cap(a))
		a = append(a, 1)
		fmt.Println("len: ", len(a), "val:", a)
	}
}
