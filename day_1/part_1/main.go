package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("day_1/input.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	calorieCount := calculateMostCalories(lines)
	fmt.Println(calorieCount)
}

func calculateMostCalories(lines []string) int {
	var mostCalories int = 0
	var currentCount int = 0

	for _, line := range lines {
		if line == "" {
			if currentCount > mostCalories {
				mostCalories = currentCount
			}
			currentCount = 0
		} else {
			calories, _ := strconv.ParseInt(line, 10, 0)
			currentCount += int(calories)
		}
	}
	if currentCount > mostCalories {
		mostCalories = currentCount
	}
	return mostCalories
}
