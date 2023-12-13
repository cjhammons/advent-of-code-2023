package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get Schematic
	schematic, err := getInput("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	sum := 0
	// Iterate the schematic and find adjacent symbols
	startIndex := -1
	endIndex := -1
	for i, row := range schematic {
		for j, n := range row {
			isInteger := isInt(n)
			if isInteger && startIndex < 0 {
				startIndex = j
			} else if isInteger && endIndex < 0 {
				endIndex = j
				if isValidPart(i, startIndex, endIndex, schematic) {
					sum += concatIndexesToInt(startIndex, endIndex, row)
				}
				startIndex = -1
				endIndex = -1
			}
		}
	}
}

func concatIndexesToInt(startIndex int, endIndex int, row []rune) int {
	var result int
	for i := startIndex; i < endIndex; i++ {
		num := row[i]
		digit, err := strconv.Atoi(string(num))
		if err != nil {
			fmt.Println("Error concating indexes to integer", err)
			return 0
		}
		result *= 10
		result += digit
	}
	return result
}

func isValidPart(rowNumber int, startIndex int, endIndex int, schematic [][]rune) bool {
	//todo
	for i := startIndex - 1; i < endIndex+1; i++ {
		// boundary checks
		if i < 0 {
			continue
		}
		if i > len(schematic[rowNumber]) {
			break
		}
		for j := rowNumber - 1; j < rowNumber+2; j++ {
			if j < 0 {
				continue
			}
			if j > len(schematic) {
				break
			}
			if isSymbol(schematic[i][j]) {
				return true
			}
		}
	}
	return false
}

func isInt(r rune) bool {
	_, err := strconv.Atoi(string(r))
	return err == nil
}

func isSymbol(r rune) bool {
	if isInt(r) || r == '.' {
		return true
	}
	return false
}

func getInput(fileName string) ([][]rune, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		result = append(result, runes)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, err
}
