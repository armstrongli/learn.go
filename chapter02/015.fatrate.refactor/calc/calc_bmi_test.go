package calculator

import (
	"testing"
)

func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.0, 1.0
	expectedOutput := 1.0
	t.Logf("开始计算，输入：height: %f, weight: %f, 期望结果：%f", inputHeight, inputWeight, expectedOutput)
	actualOutput, err := CalcBMI(inputHeight, inputWeight)
	t.Logf("实际得到：%f, error: %v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got: %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}
