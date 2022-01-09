package main

import (
	"fmt"
)

type WoodDoor struct{}

func (*WoodDoor) Open() {
	fmt.Println("WoodDoor Open")
}
func (*WoodDoor) Close() {
	fmt.Println("WoodDoor Close")
}
