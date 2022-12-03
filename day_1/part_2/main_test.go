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
	expectedCalorieCount := 45000

	res := calculateTopThreeCalories(lines)
	if res != expectedCalorieCount {
		t.Errorf("calculateTopThreeCalories: Got %d instead of %d", res, expectedCalorieCount)
	}
}
