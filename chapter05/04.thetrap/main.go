package main

import (
	"fmt"
)

func main() {
	security := Assets{assets: []Asset{
		&GlassDoor{},
		&WoodDoor{},
	}}
	fmt.Println("开始上班")
	security.DoStartWork()
	fmt.Println("8小时候，准备下班")
	security.DoStopWork()
	fmt.Println("DONE")
}
