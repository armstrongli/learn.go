package main

import (
	"fmt"
)

type GlassDoor struct{}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}
func (*GlassDoor) Close() {
	fmt.Println("GlassDoor Close")
}
