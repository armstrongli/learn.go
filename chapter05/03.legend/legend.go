package main

import (
	"fmt"
)

func legendary(legend PutElephantIntoRefrigerator, r Refrigerator, e Elephant) {
	fmt.Println("传说中，装大象可以这么装：")

	// todo human 给警告
	if _, ok := legend.(*manLegend); ok {
		fmt.Println("WARNING：现在还在用人工，效率太低")
	}

	legend.OpenTheDoorOfRefrigerator(r)
	legend.PutElephantIntoRefrigerator(e, r)
	legend.CloseTheDoorOfRefrigerator(r)

	fmt.Println("this is a legendary")
}
