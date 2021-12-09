package main

import (
	"fmt"
)

func main() {
	var p1x, p1y, p2x, p2y, p3x, p3y, p4x, p4y float64
	fmt.Print("录入第1个点的x：")
	fmt.Scanln(&p1x)
	fmt.Print("录入第1个点的y：")
	fmt.Scanln(&p1y)

	fmt.Print("录入第2个点的x：")
	fmt.Scanln(&p2x)
	fmt.Print("录入第2个点的y：")
	fmt.Scanln(&p2x)

	fmt.Print("录入第3个点的x：")
	fmt.Scanln(&p3x)
	fmt.Print("录入第3个点的y：")
	fmt.Scanln(&p3x)

	fmt.Print("录入第4个点的x：")
	fmt.Scanln(&p4x)
	fmt.Print("录入第4个点的y：")
	fmt.Scanln(&p4x)

	k1 := (p2y - p1y) / (p2x - p1x)
	k2 := (p4y - p3y) / (p4x - p3x)

	if k1 == k2 {
		fmt.Println("两条线平行")
	} else {
		fmt.Println("两条线不平行")
	}
}
