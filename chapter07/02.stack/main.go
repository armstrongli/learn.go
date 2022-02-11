package main

import (
	"fmt"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.data = append([]interface{}{data}, s.data...)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.data) > 0 {
		o := s.data[0]
		s.data = s.data[1:]
		return o, true
	}
	return nil, false
}

func main() {

	s := &Stack{}
	s.Push(111)
	s.Push(222)
	s.Push(333)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// Add(1, 2) // 展示 stack overflow
}

// func Add(a, b int) {
// 	Add(a, b)
// }
