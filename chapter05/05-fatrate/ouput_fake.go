package main

type fakeOutput struct {
	p Person
	s string
}

func (o *fakeOutput) Output(p Person, s string) {
	o.p = p
	o.s = s
}
