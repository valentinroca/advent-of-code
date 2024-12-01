package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateTotalDistance(leftNums, rightNums []int) int {
	left := make([]int, len(leftNums))
	right := make([]int, len(rightNums))
	copy(left, leftNums)
	copy(right, rightNums)

	sort.Ints(left)
	sort.Ints(right)

	totalDiff := 0
	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			totalDiff += left[i] - right[i]
		} else {
			totalDiff += right[i] - left[i]
		}
	}
	return totalDiff
}

func calculateSimilarityScore(leftNums, rightNums []int) int {
	rightCounts := make(map[int]int)
	for _, num := range rightNums {
		rightCounts[num]++
	}

	totalScore := 0
	for _, leftNum := range leftNums {
		totalScore += leftNum * rightCounts[leftNum]
	}
	return totalScore
}

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var leftNums, rightNums []int

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if len(numbers) != 2 {
			continue
		}
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])

		leftNums = append(leftNums, num1)
		rightNums = append(rightNums, num2)
	}

	totalDistance := calculateTotalDistance(leftNums, rightNums)
	similarityScore := calculateSimilarityScore(leftNums, rightNums)

	fmt.Printf("Total distance: %d\n", totalDistance)
	fmt.Printf("Similarity score: %d\n", similarityScore)
}
