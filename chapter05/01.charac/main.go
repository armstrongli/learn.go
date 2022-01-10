package main

import (
	"fmt"
	"reflect"
)

func main() {
	var refrigerator Refrigerator
	fmt.Println(refrigerator.Size)
	var elephant Elephant
	fmt.Println(elephant.Name)

	// var ip *int      // 默认是nil
	// fmt.Println(*ip) // panic 空指针
	// m := map[string]string{} // map 可以直接初始化实体，但是接口不行

	// var putER PutElephantIntoRefrigerator // 默认是nil
	// putER = PutElephantIntoRefrigerator{}
	// putER.OpenTheDoorOfRefrigerator(refrigerator) // panic 空指针

	// **注意：如果某个对象的成员方法有的在对象上，有的在对象指针上，那么在给接口赋值时，必须用指针。**
	var legend PutElephantIntoRefrigerator = &PutElephantIntoRefrigeratorImpl{}
	legend.OpenTheDoorOfRefrigerator(refrigerator)
	legend.PutElephantIntoRefrigerator(elephant, refrigerator)
	legend.CloseTheDoorOfRefrigerator(refrigerator)

	// var o Open = Refrigerator{}
	var c Close = Refrigerator{}
	var b Box = Refrigerator{}
	fmt.Println(b, c)
	c = b // 从范围小的向范围大的转会成功
	// b = c // 从范围大的向范围小的转会失败

	var i interface{}
	i = 3
	fmt.Println(reflect.TypeOf(i), "value:", i)
	i = 3.2345
	fmt.Println(reflect.TypeOf(i), "value:", i)
	i = Refrigerator{}
	fmt.Println(reflect.TypeOf(i), "value:", i)
}

type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant, Refrigerator) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
}

type TestTypeImplInterface func()

func (t TestTypeImplInterface) OpenTheDoorOfRefrigerator(_ Refrigerator) error {
	return nil
}
func (t TestTypeImplInterface) PutElephantIntoRefrigerator(_ Elephant, _ Refrigerator) error {
	return nil
}
func (t TestTypeImplInterface) CloseTheDoorOfRefrigerator(_ Refrigerator) error {
	return nil
}

type PutElephantIntoRefrigeratorImpl struct {
}

func (legent *PutElephantIntoRefrigeratorImpl) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("打开冰箱门")
	return nil
}

func (legent *PutElephantIntoRefrigeratorImpl) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	// todo
	fmt.Println("装进去")
	return nil
}

func (legent *PutElephantIntoRefrigeratorImpl) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("关门")
	return nil
}

type Open interface {
	Open() error
}

type Close interface {
	Close() error
}

type Box interface {
	Open
	Close
}

type Refrigerator struct {
	Size string
}

func (Refrigerator) Open() error {
	return nil
}

func (Refrigerator) Close() error {
	return nil
}

type Elephant struct {
	Name string
}
