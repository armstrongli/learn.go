package main

import (
	"log"

	gobmi "github.com/armstrongli/go-bmi"
)

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int

	bmi     float64
	fatRate float64
}

func (p *Person) calcBmi() error {
	bmi, err := gobmi.BMI(p.weight, p.tall)
	if err != nil {
		log.Printf("error when calculating BMP for Person[%s]: %v", p.name, err)
		return err
	}
	p.bmi = bmi
	return nil
}

func (p *Person) calcFatRate() {
	p.fatRate = gobmi.CalcFatRate(p.bmi, p.age, p.sex)
}
