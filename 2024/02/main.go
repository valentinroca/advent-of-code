package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	isIncreasing := levels[1] > levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if isIncreasing && diff < 0 {
			return false
		}
		if !isIncreasing && diff > 0 {
			return false
		}
	}

	return true
}

func isReportSafeWithDampener(levels []int) bool {
	if isReportSafe(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		dampened := make([]int, 0, len(levels)-1)
		dampened = append(dampened, levels[:i]...)
		dampened = append(dampened, levels[i+1:]...)

		if isReportSafe(dampened) {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0
	safeCountWithDampener := 0

	for scanner.Scan() {
		levels := make([]int, 0)
		for _, field := range strings.Fields(scanner.Text()) {
			if num, err := strconv.Atoi(field); err == nil {
				levels = append(levels, num)
			}
		}

		if isReportSafe(levels) {
			safeCount++
			safeCountWithDampener++
		} else if isReportSafeWithDampener(levels) {
			safeCountWithDampener++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
	fmt.Printf("Number of safe reports with dampener: %d\n", safeCountWithDampener)
}
