package main

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	r := TestBox{}
	var c Close = r

	switch cDetail := c.(type) {
	case Refrigerator:
		fmt.Println("是预期的冰箱")
		fmt.Println(cDetail.Size)
	case TestBox:
		fmt.Println("这是一个box，不能当冰箱用")
	}
	{
		refrigerator, ok := checkIsRefrigerator(c)
		if ok {
			fmt.Println("是个冰箱，开门装大象", refrigerator)
		} else {
			fmt.Println("这不是个冰箱")
		}
	}
}

func checkIsRefrigerator(c Close) (Refrigerator, bool) {
	r, ok := c.(Refrigerator)
	return r, ok
}

type TestBox struct {
}

func (tb TestBox) Close() error {
	return nil
}
