package main

type InputService interface {
	GetInput() Person
}

type OutputService interface {
	Output(Person, string)
}
