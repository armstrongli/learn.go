package main

type fakeInput struct {
}

func (*fakeInput) GetInput() Person {
	return Person{
		name:   "小强",
		sex:    "男",
		tall:   1.7,
		weight: 70,
		age:    35,
	}
}
