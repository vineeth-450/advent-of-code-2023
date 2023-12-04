package main

import (
	"bufio"
	"fmt"
	"os"
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

		numbers := strings.Split(line, ":")[1]

		winningNumbersStr := strings.TrimSpace(strings.Split(numbers, "|")[0])
		winningNumbers := strings.Split(winningNumbersStr, " ")
		myNumbersStr := strings.TrimSpace(strings.Split(numbers, "|")[1])
		myNumbers := strings.Split(myNumbersStr, " ")

		winningNumbersMap := make(map[string]struct{})
		for _, n := range winningNumbers {
			winningNumbersMap[n] = struct{}{}
		}

		cardWorth := 0
		for _, n := range myNumbers {
			if n == "" {
				continue
			}
			if _, ok := winningNumbersMap[n]; ok {

				if cardWorth == 0 {
					cardWorth = 1
				} else {
					cardWorth *= 2
				}

			}
		}

		sum += cardWorth
	}

	fmt.Println("Answer for part one is", sum)
}

func partTwo() {
	filename := "./input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	lineCount := 0
	copiesMap := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_ = scanner.Text()
		lineCount += 1
		copiesMap[lineCount] = 1
	}

	file, err = os.Open(filename)
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	lineCount = 0
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount += 1
		line := scanner.Text()

		numbers := strings.Split(line, ":")[1]

		winningNumbersStr := strings.TrimSpace(strings.Split(numbers, "|")[0])
		winningNumbers := strings.Split(winningNumbersStr, " ")
		myNumbersStr := strings.TrimSpace(strings.Split(numbers, "|")[1])
		myNumbers := strings.Split(myNumbersStr, " ")

		winningNumbersMap := make(map[string]struct{})
		for _, n := range winningNumbers {
			winningNumbersMap[n] = struct{}{}
		}

		matchCount := 0
		for _, n := range myNumbers {
			if n == "" {
				continue
			}
			if _, ok := winningNumbersMap[n]; ok {
				matchCount += 1
			}
		}

		for i := lineCount + 1; i <= lineCount+matchCount; i++ {
			if _, ok := copiesMap[i]; ok {
				copiesMap[i] = copiesMap[i] + copiesMap[lineCount]
			}
		}
	}

	sum := 0
	for _, v := range copiesMap {
		sum += v
	}

	fmt.Println("Answer for part two is", sum)
}
