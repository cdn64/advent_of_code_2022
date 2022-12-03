package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("day_1/input.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	calorieCount := calculateTopThreeCalories(lines)
	fmt.Println(calorieCount)
}

func calculateTopThreeCalories(lines []string) int {
	calorieCounts := []int{}
	var currentCount int = 0

	for _, line := range lines {
		if line == "" {
			calorieCounts = append(calorieCounts, currentCount)
			currentCount = 0
		} else {
			calories, _ := strconv.ParseInt(line, 10, 0)
			currentCount += int(calories)
		}
	}
	calorieCounts = append(calorieCounts, currentCount)

	sort.Sort(sort.Reverse(sort.IntSlice(calorieCounts)))
	calorieCounts = calorieCounts[0:3]

	total := 0
	for _, calorieCount := range calorieCounts {
		total += calorieCount
	}

	return total
}
