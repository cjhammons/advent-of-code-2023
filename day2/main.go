package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const RED_COUNT_MAX = 12
const GREEN_COUNT_MAX = 13
const BLUE_COUNT_MAX = 14

type Game struct {
	Id    int
	Red   int
	Green int
	Blue  int
}

func main() {
	input, err := getInput("input.txt")
	if err != nil {
		os.Exit(1)
	}

	id_sum := 0

	for _, line := range input {
		game, isValid := parseGame(line)
		if isValid {
			println("Valid game: ", game.Id)
			println("Red: ", game.Red)
			println("Blue ", game.Blue)
			println("Green: ", game.Green)
			id_sum += game.Id
			println("Sum is currently: ", id_sum)
		} else {
			println("InValid  game: ", game.Id)
			println("Red: ", game.Red)
			println("Blue ", game.Blue)
			println("Green: ", game.Green)
			println("Sum remains: ", id_sum)
		}
	}

	println("Id sum is: ", id_sum)
}

func isValidDraw(r, b, g int) bool {
	return r <= RED_COUNT_MAX && b <= BLUE_COUNT_MAX && g <= GREEN_COUNT_MAX
}

func parseGame(line string) (Game, bool) {
	parts := strings.Split(line, ": ")
	id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])

	draws := strings.Split(parts[1], "; ")

	game := Game{
		Id:    id,
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for _, d := range draws {
		r, b, g := parseDraw(d)
		if !isValidDraw(r, b, g) {
			return Game{}, false
		}
		game.Red += r
		game.Blue += b
		game.Green += g
	}

	return game, true
}

func parseDraw(draw string) (int, int, int) {
	redCount := 0
	greenCount := 0
	blueCount := 0

	// Compile the regular expression
	re := regexp.MustCompile(`(\d+) (red|green|blue)`)
	matches := re.FindAllStringSubmatch(draw, -1)

	for _, match := range matches {
		count, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error converting count:", err)
			os.Exit(1)
		}
		color := match[2]

		switch color {
		case "blue":
			blueCount += count
		case "red":
			redCount += count
		case "green":
			greenCount += count
		}
	}

	return redCount, blueCount, greenCount
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
