package main

import (
	"testing"
)

func TestCase1Part1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期 王强 第一，但是得到的是：%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期 王强的体脂是 0.32，但得到的是：%f", fatRateOfWQ)
		}
	}
}

func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期 王强 第一，但是得到的是：%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期 王强的体脂是 0.32，但得到的是：%f", fatRateOfWQ)
		}
	}
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期 王强 第一，但是得到的是：%d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期 王强的体脂是 0.32，但得到的是：%f", fatRateOfWQ)
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是：%d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期 李静的体脂是 0.28，但得到的是：%f", fatRateOfLJ)
		}
	}
}

func TestCase2(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("张伟", 0.38)
	inputRecord("李静", 0.28)

	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是：%d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期 李静的体脂是 0.28，但得到的是：%f", fatRateOfLJ)
		}
	}
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期 王强 第2，但是得到的是：%d", randOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期 王强的体脂是 0.38，但得到的是：%f", fatRateOfWQ)
		}
	}
	{
		randOfZW, fatRateOfZW := getRand("张伟")
		if randOfZW != 2 {
			t.Fatalf("预期 张伟 第2，但是得到的是：%d", randOfZW)
		}
		if fatRateOfZW != 0.38 {
			t.Fatalf("预期 张伟的体脂是 0.38，但得到的是：%f", fatRateOfZW)
		}
	}
}

func TestCase3(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("李静", 0.28)
	inputRecord("张伟")

	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是：%d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期 李静的体脂是 0.28，但得到的是：%f", fatRateOfLJ)
		}
	}
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期 王强 第一，但是得到的是：%d", randOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期 王强的体脂是 0.38，但得到的是：%f", fatRateOfWQ)
		}
	}
	{
		randOfZW, _ := getRand("张伟")
		if randOfZW != 3 {
			t.Fatalf("预期 张伟 第三，但是得到的是：%d", randOfZW)
		}
	}
}
