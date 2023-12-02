package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := 0, 0
		firstDigitSet := false
		for _, c := range line {
			d, err := strconv.Atoi(string(c))
			if err == nil && !firstDigitSet {
				firstDigit, lastDigit = d, d
				firstDigitSet = true
			} else if err == nil {
				lastDigit = d
			}
		}

		numberStr := fmt.Sprintf("%d%d", firstDigit, lastDigit)
		d, err := strconv.Atoi(string(numberStr))
		if err != nil {
			fmt.Println("Error converting formed 2 digit number", err)
			return
		}

		sum += d
	}

	fmt.Println("Answer for part one is", sum)
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		d := getTwoDigitNumberPartTwo(line)

		sum += d
	}

	fmt.Println("Answer for part two is", sum)
}

func getTwoDigitNumberPartTwo(line string) int {
	firstDigit, lastDigit := 0, 0
	firstDigitSet := false

	ind, nextDigit := 0, 0
	var err error
	for ind < len(line) {
		nextDigit, ind, err = getNextDigit(line, ind)

		if err == nil && !firstDigitSet {
			firstDigit, lastDigit = nextDigit, nextDigit
			firstDigitSet = true
		} else if err == nil {
			lastDigit = nextDigit
		}
	}

	numberStr := fmt.Sprintf("%d%d", firstDigit, lastDigit)
	d, err := strconv.Atoi(string(numberStr))
	if err != nil {
		fmt.Println("Error converting formed 2 digit number", err)
		return 0
	}

	return d
}

var digits = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func getNextDigit(str string, ind int) (int, int, error) {

	for i := ind; i < len(str); i++ {
		d, err := strconv.Atoi(string(str[i]))
		if err == nil {
			return d, i + 1, nil
		}
		for k, v := range digits {
			digitStrLen := len(k)
			indexEnd := i + digitStrLen
			if indexEnd >= len(str) {
				indexEnd = len(str)
			}

			if k == str[i:indexEnd] {
				return v, i + 1, nil
			}
		}
	}

	return 0, ind + 1, errors.New("no digit Found")
}
