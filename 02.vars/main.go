package main

import (
	"fmt"
	"math"
)

func main() {
	var hello string = "hello, golang!"
	var world = "world"
	fmt.Println(hello, world)
	小数 := 1.234
	fmt.Println(小数)

	var int3, int4 uint = 33, 44
	fmt.Println(int3 * int4)

	 ho, ver  := 3, 4.123
	var sc = ho * int(ver)
	fmt.Println(ho * int(ver))
	fmt.Println(sc)
	var newname string
	fmt.Println("newname=", newname)

	var int6 uint = math.MaxUint64
	fmt.Println(int6)
	var int7 int = int(int6)
	fmt.Println(int7)

	// var nameOfSquare string
	// var _name string
}
