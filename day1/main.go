package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Define a map for string to integer conversion
var stringToInt = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func main() {
	input, err := getInput("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Part One...")
	nums := parseLines(input)
	sum := sumNums(nums)

	fmt.Println("Answer is: ", sum)
}

func getDigits(line string) int {
	var first rune
	var subString []rune
	for _, c := range line {
		if unicode.IsDigit(c) {
			first = c
			break
		} else {
			subString = append(subString, c)
			if isNumString(string(subString)) {
				first = convertStringToInt(string(subString))
				break
			}
			if !isSubsetOfAny(string(subString)) {
				subString = subString[:0]
			}	
		}
	}

	var last rune
	line_runes := []rune(line)
	for i := len(line_runes) - 1; i >= 0; i-- {
		c := line_runes[i]
		if unicode.IsDigit(c) {
			last = c
			break
		} else {
			subString = append(subString, c)
			if isNumString(string(subString)) {
				last = convertStringToInt(string(subString))
				break
			}
			if !isSubsetOfAny(string(subString)) {
				subString = subString[:0]
			}	
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

// Function to convert a string to its corresponding integer
func convertStringToInt(s string) rune {
	val := stringToInt[s]
	return val
}

func isNumString(s string) (bool) {
	for key := range stringToInt{
		if s == key {
			return true
		}
	}
	return false
}

// returns:
//
//	bool: If input is a subset
//	bool: If input == a string number
func isSubsetOfAny(s string) (bool) {
	for key := range stringToInt {
		if s == key {
			return true
		} else if strings.Contains(key, s) {
			return true
		}
	}
	return false
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
