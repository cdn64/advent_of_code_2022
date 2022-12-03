package main

import "testing"

func TestExample(t *testing.T) {
	lines := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
	expectedCalorieCount := 24000

	res := calculateMostCalories(lines)
	if res != expectedCalorieCount {
		t.Errorf("calculateMostCalories: Got %d instead of %d", res, expectedCalorieCount)
	}
}
