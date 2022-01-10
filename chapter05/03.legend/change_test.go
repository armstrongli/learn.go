package main

import (
	"fmt"
	"testing"
)

type Change interface {
	ChangeName(newName string)
	ChangeAge(newAge int)
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) ChangeName(newName string) {
	s.Name = newName
}

func (s Student) ChangeAge(newAge int) {
	s.Age = newAge
}

func TestVal(t *testing.T) {
	var stdChg Change
	// stdChg = Student{} // 当实现接口的方法是在对象指针上时，只能用对象指针作为值赋值给接口
	stdChg = &Student{
		Name: "Tom",
		Age:  0,
	}

	fmt.Println(stdChg)
}

// func TestStd(t *testing.T) {
// 	s := Student{Name: "Tom"}
// 	fmt.Println(s.Name)
// 	s.ChangeName("Jerry")
// 	fmt.Println(s.Name)
// }
