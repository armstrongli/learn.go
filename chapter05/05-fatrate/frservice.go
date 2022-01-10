package main

import (
	"log"
)

type fatRateService struct {
	input  InputService
	s      *fatRateSuggestion
	output OutputService
}

func (frsvc *fatRateService) GiveSuggestionToPerson(person *Person) {
	if err := person.calcBmi(); err != nil {
		log.Printf("无法给%s计算BMI:%v", person.name, err)
		return
	}
	person.calcFatRate()
	frsvc.output.Output(*person, frsvc.s.GetSuggestion(person))
}

// func (frsvc *fatRateService) GiveSuggestionToPersons(persons ...*Person) map[*Person]string {
// 	out := map[*Person]string{}
// 	for _, item := range persons {
// 		out[item] = frsvc.GiveSuggestionToPerson(item)
// 	}
// 	return out
// }
