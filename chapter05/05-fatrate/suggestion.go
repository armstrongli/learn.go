package main

func GetFatRateSuggestion() *fatRateSuggestion {
	return &fatRateSuggestion{
		suggArr: [][][]int{
			{ // 第一个元素，表示男
				{ // 18-39
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				},
				{ // 40-59

				},
				{ // 60+

				},
			},
			{ // 第二个元素表示女
				{ // 18-39

				},
				{ // 40-59

				},
				{ // 60+

				},
			},
		},
	}
}

type fatRateSuggestion struct {
	suggArr [][][]int
}

func (s *fatRateSuggestion) GetSuggestion(person *Person) string {
	sexIdx := s.getIndexOfSex(person.sex)
	ageIdx := s.getIndexOfAge(person.age)

	maxFRSupported := len(s.suggArr[sexIdx][ageIdx]) - 1
	frIdx := int(person.fatRate * 100)
	if frIdx > maxFRSupported {
		frIdx = maxFRSupported
	}
	suggIdx := s.suggArr[sexIdx][ageIdx][frIdx]
	return s.translateResult(suggIdx)
}

func (s *fatRateSuggestion) getIndexOfSex(sex string) int {
	if sex == "男" {
		return 0
	}
	return 1
}

func (s *fatRateSuggestion) getIndexOfAge(age int) int {
	switch {
	case age >= 18 && age <= 39:
		return 0
	case age >= 40 && age <= 59:
		return 1
	case age >= 60:
		return 2
	default:
		return -1
	}
}
func (s *fatRateSuggestion) translateResult(idx int) string {
	switch idx {
	case 0:
		return "偏瘦"
	case 1:
		return "标准"
	case 2:
		return "偏重"
	case 3:
		return "肥胖"
	case 4:
		return "非常肥胖"
	}
	return "未知"
}
