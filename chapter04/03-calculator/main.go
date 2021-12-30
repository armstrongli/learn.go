package main

import (
	"fmt"
)

func main() {
	var left, right int = 1, 2
	// var op string = "+"

	c := Calculator{
		left:  left,
		right: right,
		// op:    op,
	}
	fmt.Printf("&c @ main =%p\n", &c)
	fmt.Println(c.Add())
	fmt.Println("c.result=", c.result)

	newC := NewCalculator{}
	newC.left = 100
	newC.right = 200
	fmt.Println(newC.Add())

	mc := MyCommand{}
	mc.commandOptions["aa"] = "AAA"
	fmt.Println(mc.ToCmdStr())
}

type MyCommand struct {
	mainCommand    *string
	commandOptions map[string]string
}

func (my MyCommand) ToCmdStr() string {
	out := ""
	for k, v := range my.commandOptions {
		out = out + fmt.Sprintf("--%s=%s", k, v)
	}
	return out
}
