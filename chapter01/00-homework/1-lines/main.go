package main

import (
	"fmt"
)

func main() {
	var p1x, p1y, p2x, p2y, p3x, p3y, p4x, p4y float64
	p1x, p1y = getPointXYFromInput()
	p2x, p2y = getPointXYFromInput()
	p3x, p3y = getPointXYFromInput()
	p4x, p4y = getPointXYFromInput()

	k1 := calculateK(p2y, p1y, p2x, p1x)
	k2 := calculateK(p4y, p3y, p4x, p3x)

	getResult(k1, k2)
}

func getPointXYFromInput() (float64, float64) {
	var x, y float64
	fmt.Print("录入x：")
	fmt.Scanln(&x)
	fmt.Print("录入y：")
	fmt.Scanln(&y)
	return x, y
}

func calculateK(y2, y1, x2, x1 float64) float64 {
	return (y2 - y1) / (x2 - x1)
}

func getResult(k1, k2 float64) {
	if k1 == k2 {
		fmt.Println("两条线平行")
	} else {
		fmt.Println("两条线不平行")
	}
}
