package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func processLine(line string, re *regexp.Regexp) int {
	lineSum := 0
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		lineSum += x * y
	}
	return lineSum
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	scanner := bufio.NewScanner(file)
	mulSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		mulSum += processLine(line, re)
	}

	fmt.Printf("Sum of all multiplication results: %d\n", mulSum)
}
