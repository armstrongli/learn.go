package main

import (
	"fmt"
)

type StdOut struct {
}

func (*StdOut) Output(p Person, s string) {
	fmt.Println(s)
}
