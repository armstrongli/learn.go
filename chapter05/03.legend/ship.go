package main

import (
	"fmt"
)

type shipLegend struct {
}

func (*shipLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 ship 做 OpenTheDoorOfRefrigerator")
	return nil
}
func (*shipLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 ship 做 PutElephantIntoRefrigerator")
	return nil
}
func (*shipLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 ship 做 CloseTheDoorOfRefrigerator")
	return nil
}
