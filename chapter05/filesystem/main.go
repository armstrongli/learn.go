package main

import (
	"fmt"
)

func main() {
	var data string
	{
		var equipment IOInterface = &Soft{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Sata{}
		data = equipment.Read()
		fmt.Println(data)
	}
}

type IOInterface interface {
	Read() (data string)
}

type Sata struct{}

func (Sata) Read() string {
	return "安安静静的sata"
}

type Soft struct {
}

func (Soft) Read() string {
	return "软盘数据，啦啦啦啦"
}

type Mag struct {
}

func (Mag) Read() string {
	return "滋滋滋滋磁带"
}

type Paper struct {
}

func (Paper) Read() string {
	return "从纸带读*******00000"
}
