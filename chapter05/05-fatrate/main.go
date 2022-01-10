package main

func main() {
	frSvc := &fatRateService{
		s:      GetFatRateSuggestion(),
		input:  &inputFromStd{},
		output: &StdOut{},
	} // todo

	for {
		p := frSvc.input.GetInput()
		frSvc.GiveSuggestionToPerson(&p)
	}
}
