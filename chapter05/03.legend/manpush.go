package main

import (
	"fmt"
)

type manLegend struct {
}

func (*manLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manLegend 做 OpenTheDoorOfRefrigerator")
	return nil
}
func (*manLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 manLegend 做 PutElephantIntoRefrigerator")
	return nil
}
func (*manLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manLegend 做 CloseTheDoorOfRefrigerator")
	return nil
}
