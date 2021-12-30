package main

import (
	"log"
)

type fatRateService struct {
	s *fatRateSuggestion
}

func (frsvc *fatRateService) GiveSuggestionToPerson(person *Person) string {
	if err := person.calcBmi(); err != nil {
		log.Printf("无法给%s计算BMI:%v", person.name, err)
		return "错误"
	}
	person.calcFatRate()
	return frsvc.s.GetSuggestion(person)
}

func (frsvc *fatRateService) GiveSuggestionToPersons(persons ...*Person) map[*Person]string {
	out := map[*Person]string{}
	for _, item := range persons {
		out[item] = frsvc.GiveSuggestionToPerson(item)
	}
	return out
}
