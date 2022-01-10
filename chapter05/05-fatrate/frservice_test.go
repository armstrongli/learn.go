package main

import (
	"testing"
)

func TestFrServiceSuggestion(t *testing.T) {
	realOutput := &fakeOutput{}
	frSvc := &fatRateService{
		s:      GetFatRateSuggestion(),
		input:  &fakeInput{},
		output: realOutput,
	}
	p := frSvc.input.GetInput()
	expOutput := &fakeOutput{
		p: p,
		s: "偏重",
	}
	frSvc.GiveSuggestionToPerson(&p)

	if realOutput.s != expOutput.s {
		t.Fatalf("预期：%s, 得到：%s", expOutput.s, realOutput.s)
	}
}
