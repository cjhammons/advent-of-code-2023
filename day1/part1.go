package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func partOne() {
	input, err := getInput("input.txt")
	if err != nil {
		os.Exit(1)
	}
	nums := parseLines(input)
	sum := sumNums(nums)

	fmt.Println("Answer is: ", sum)
}

func getInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getDigits(line string) int {
	var first rune
	for _, c := range line {
		if unicode.IsDigit(c) {
			first = c
			break
		}
	}

	var last rune
	line_runes := []rune(line)
	for i := len(line_runes) - 1; i >= 0; i-- {
		c := line_runes[i]
		if unicode.IsDigit(c) {
			last = c
			break
		}
	}

	concat := string(first) + string(last)
	intValue, err := strconv.Atoi(concat)
	if err != nil {
		fmt.Println("Error converting to int:", err)
		os.Exit(1)
	}

	return intValue
}

func parseLines(lines []string) []int {
	var nums []int
	for _, line := range lines {
		digits := getDigits(line)
		nums = append(nums, digits)
		fmt.Printf("Line %s \nPoduced digits %d\n\n", line, digits)
	}
	fmt.Println("Numbers: ", nums)
	return nums
}

func sumNums(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	fmt.Println("Sum is: ", sum)
	return sum
}
