package main

import (
	"fmt"

	gobmi "github.com/armstrongli/go-bmi"
)

func main() {
	// name1, sex1, tall1, weight1, age1 := // ...
	// name2, sex2, tall2, weight2, age2 := // ...
	// name3, sex3, tall3, weight3, age3 := // ...

	// names := []string{}
	// sexes := []string{}
	// talls := []float64{}
	// weights := []float64{}
	// ages := []int{}

	persons := []Person{
		{
			"小强",
			"男",
			1.7,
			70,
			35,
		},
	}
	xq := Person{
		"小强",
		"男",
		1.7,
		70,
		35,
	}
	fmt.Println(xq)
	for _, item := range persons {
		bmi, err := gobmi.BMI(item.weight, item.tall)
		fmt.Println(bmi, err)
	}
	// a:=new(Person)
	// b:=&Person{}
}

type Person struct {
	name string
	// Alias []string
	sex    string
	tall   float64
	weight float64
	age    int
}
