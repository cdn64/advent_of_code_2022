package main

import "testing"

func TestExample(t *testing.T) {
	lines := []string{
		"A Y",
		"B X",
		"C Z",
	}
	expectedTotalScore := 15

	res := calculateExpectedTotalScore(lines)
	if res != expectedTotalScore {
		t.Errorf("calculateExpectedTotalScore: Got %d instead of %d", res, expectedTotalScore)
	}
}
