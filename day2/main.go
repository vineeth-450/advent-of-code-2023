package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ":")
		game := split[0]

		gameNoStr := ""
		if len(game) == 7 {
			gameNoStr = game[5:7]
		} else if len(game) == 8 {
			gameNoStr = game[5:8]
		} else {
			gameNoStr = game[5:6]
		}

		gameNo, err := strconv.Atoi(string(gameNoStr))
		if err != nil {
			fmt.Println("Error converting game number to int", err)
			return
		}

		outputs := split[1]
		sets := strings.Split(outputs, ";")
		if isGamePossible(sets) {
			sum += gameNo
		}
	}

	fmt.Println("Answer for part one is", sum)
}

// 12 red cubes, 13 green cubes, and 14 blue cubes
var cubesLimit = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isGamePossible(sets []string) bool {
	for _, set := range sets {
		cubes := strings.Split(set, ",")

		for _, colorCount := range cubes {
			colorCount = strings.TrimSpace(colorCount)
			s := strings.Split(colorCount, " ")
			count := s[0]
			countNo, err := strconv.Atoi(count)
			if err != nil {
				fmt.Println("Error while coverting count to int", err)
				return false
			}
			color := s[1]
			if countNo > cubesLimit[color] {
				return false
			}
		}
	}

	return true
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ":")
		outputs := split[1]
		sets := strings.Split(outputs, ";")

		sum += getGamePower(sets)
	}

	fmt.Println("Answer for part two is", sum)
}

func getGamePower(sets []string) int {
	maxSet := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range sets {
		cubes := strings.Split(set, ",")

		for _, colorCount := range cubes {
			colorCount = strings.TrimSpace(colorCount)
			s := strings.Split(colorCount, " ")
			count := s[0]
			color := s[1]

			countNo, err := strconv.Atoi(count)
			if err != nil {
				fmt.Println("Error while coverting count to int", err)
				return 0
			}

			if countNo > maxSet[color] {
				maxSet[color] = countNo
			}
		}
	}

	return maxSet["green"] * maxSet["blue"] * maxSet["red"]
}
