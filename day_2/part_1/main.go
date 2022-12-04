package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("day_2/input.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	expectedTotalScore := calculateExpectedTotalScore(lines)
	fmt.Println(expectedTotalScore)
}

const COLUMN_1_BASE = 65
const COLUMN_2_BASE = 88

func calculateExpectedTotalScore(lines []string) int {
	totalScore := 0

	for _, line := range lines {
		if len(line) != 3 {
			continue
		}

		opponent := int(line[0] - COLUMN_1_BASE)
		player := int(line[2] - COLUMN_2_BASE)
		totalScore += calculateScore(player, opponent)
		totalScore += player + 1
	}

	return totalScore
}

// 0 Rock
// 1 Paper
// 2 Scissors
func calculateScore(player int, opponent int) int {
	if player == opponent {
		return 3
	} else {
		diff := player - opponent
		if diff == 1 || diff == -2 {
			return 6
		} else {
			return 0
		}
	}
}
