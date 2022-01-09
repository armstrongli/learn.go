package main

import (
	"fmt"
)

var _ Door = &WoodDoor{}

type WoodDoor struct{}

func (d *WoodDoor) Unlock() {
	fmt.Println("WoodDoor Unlock")
}

func (d *WoodDoor) Lock() {
	fmt.Println("WoodDoor Lock")
}

func (*WoodDoor) Open() {
	fmt.Println("WoodDoor Open")
}
func (*WoodDoor) Close() {
	fmt.Println("WoodDoor Close")
}
