package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	sync.Mutex
	data []interface{}
}

func (q *Queue) Push(data interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, data)
}

func (q *Queue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.data) > 0 {
		o := q.data[0]
		q.data = q.data[1:]
		return o, true
	}
	return nil, false
}

func main() {
	q := &Queue{}
	q.Push(111)
	q.Push(222)
	q.Push(333)
	q.Push(nil)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
