package calculator

import (
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcBMI(tall float64, weight float64) (bmi float64) {
	bmi, _ = gobmi.BMI(weight, tall)
	return
}
