package main

var defaultCalculator = Calculator{}

type Calculator struct {
	left, right int
	op          string
	result      int
}
