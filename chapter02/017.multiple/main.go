package main

import (
	"fmt"
)

func main() {
	// bmis := []float64{1.2, 3.222, 4.343443}
	avgBmi := calculateAvg(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(avgBmi)
	avgBmi = calculateAvgOnSlice([]float64{})
	fmt.Println(avgBmi)
	getScoresOfStudent("Tom")
}

func calculateAvg(bmis ...float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	avgBmi = total / float64(len(bmis))
	return
}

func calculateAvgOnSlice(bmis []float64) float64 {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total / float64(len(bmis))
}

func getScoresOfStudent(stdName string) (chinese int, match int, english int, physics int, nature int) {
	return 0, 0, 0, 0, 0
}
