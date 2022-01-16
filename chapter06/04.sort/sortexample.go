package main

import (
	"fmt"
	"sort"
)

type Button struct {
	Floor int
}

type Elevator struct {
	buttons  Buttons
	position int
}

type Buttons []*Button

func (b Buttons) Len() int {
	return len(b)
}

func (b Buttons) Less(i, j int) bool {
	return b[i].Floor < b[j].Floor
}

func (b Buttons) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
	// tmpObj := b[i]
	// b[i] = b[j]
	// b[j] = tmpObj
}

func main() {
	ev := &Elevator{
		position: 2,
		buttons: Buttons{
			{Floor: 3},
			{Floor: 1},
			{Floor: 5},
			{Floor: 2},
			{Floor: 4},
		},
	}
	sort.Sort(sort.Reverse(ev.buttons))

	fmt.Printf("%+v\n", ev.buttons)
	for _, item := range ev.buttons {
		fmt.Println(item.Floor)
	}
}

// func sortButtons([]*Button)[]*Button{
// 	for
// }
