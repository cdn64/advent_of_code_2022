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
		outcome := int(line[2] - COLUMN_2_BASE)
		player := calculatePlayerMove(opponent, outcome)
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
		diff := (3 + player - opponent) % 3
		if diff == 1 {
			return 6
		} else {
			return 0
		}
	}
}

func calculatePlayerMove(opponent int, outcome int) int {
	if outcome == 0 { // Lose
		return (opponent + 2) % 3
	} else if outcome == 2 { // Win
		return (opponent + 1) % 3
	}
	return opponent // Draw
}
