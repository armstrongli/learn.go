package main

import (
	"fmt"
)

func main() {
	var refrigerator Refrigerator
	fmt.Println(refrigerator.Size)
	var elephant Elephant
	fmt.Println(elephant.Name)

	var ip *int      // 默认是nil
	fmt.Println(*ip) // panic 空指针

	var putER PutElephantIntoRefrigerator         // 默认是nil
	putER.OpenTheDoorOfRefrigerator(refrigerator) // panic 空指针
}

type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant, Refrigerator) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
}

type Refrigerator struct {
	Size string
}

type Elephant struct {
	Name string
}
