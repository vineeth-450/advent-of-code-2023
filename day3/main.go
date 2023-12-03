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

	engineSchema := [][]string{}

	lineNo := 0
	specialCharMap := map[int][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		lineSchema := []string{}

		for j, c := range line {
			if checkIfSpecialChar(string(c)) {
				specialCharMap[lineNo] = append(specialCharMap[lineNo], j)
			}
			lineSchema = append(lineSchema, string(c))
		}

		engineSchema = append(engineSchema, lineSchema)
		lineNo += 1
	}

	for i, lineSchema := range engineSchema {
		numberStr := ""
		charAround := false

		for j, item := range lineSchema {
			_, err := strconv.Atoi(string(item))
			if err == nil {
				if !charAround && checkIfSpecialCharAround(engineSchema, specialCharMap, i, j) {
					charAround = true
				}

				numberStr += item

				if j == len(lineSchema)-1 {
					number, err := strconv.Atoi(numberStr)
					if err != nil {
						fmt.Println("Error converting number", err)
					}

					if charAround {
						sum += number
						charAround = false
					}
					numberStr = ""
				}
			} else if len(numberStr) > 0 {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Println("Error converting number", err)
				}

				if charAround {
					sum += number
					charAround = false
				}
				numberStr = ""
			} else {
				numberStr = ""
			}
		}

	}

	fmt.Println("Answer for part one is", sum)
}

var nonSpecialChars = "1234567890."

func checkIfSpecialCharAround(engineSchema [][]string, specialCharMap map[int][]int, i, j int) bool {
	return checkIfEntryExist(specialCharMap, i-1, j-1) || checkIfEntryExist(specialCharMap, i, j-1) || checkIfEntryExist(specialCharMap, i+1, j-1) ||
		checkIfEntryExist(specialCharMap, i-1, j) || checkIfEntryExist(specialCharMap, i+1, j) || checkIfEntryExist(specialCharMap, i-1, j+1) ||
		checkIfEntryExist(specialCharMap, i, j+1) || checkIfEntryExist(specialCharMap, i+1, j+1)
}

func returnIfStarAround(engineSchema [][]string, specialCharMap map[int][]int, i, j int) (int, int) {
	if a, b := returnIfEntryExist(specialCharMap, i-1, j-1); a != -1 {
		return a, b
	}
	if a, b := returnIfEntryExist(specialCharMap, i, j-1); a != -1 {
		return a, b
	}
	if a, b := returnIfEntryExist(specialCharMap, i+1, j-1); a != -1 {
		return a, b
	}
	if a, b := returnIfEntryExist(specialCharMap, i-1, j); a != -1 {
		return a, b
	}
	if a, b := returnIfEntryExist(specialCharMap, i+1, j); a != -1 {
		return a, b
	}
	if a, b := returnIfEntryExist(specialCharMap, i-1, j+1); a != -1 {
		return a, b
	}
	if a, b := returnIfEntryExist(specialCharMap, i, j+1); a != -1 {
		return a, b
	}
	if a, b := returnIfEntryExist(specialCharMap, i+1, j+1); a != -1 {
		return a, b
	}

	return -1, -1
}

func checkIfSpecialChar(s string) bool {
	return !strings.Contains(nonSpecialChars, s)
}

func checkIfEntryExist(specialCharMap map[int][]int, i, j int) bool {
	for _, v := range specialCharMap[i] {
		if v == j {
			return true
		}
	}

	return false
}

func returnIfEntryExist(specialCharMap map[int][]int, i, j int) (int, int) {
	for _, v := range specialCharMap[i] {
		if v == j {
			return i, j
		}
	}

	return -1, -1
}

func checkIfStar(s string) bool {
	return s == "*"
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	sum := 0
	scanner := bufio.NewScanner(file)

	engineSchema := [][]string{}

	lineNo := 0
	specialCharMap := map[int][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		lineSchema := []string{}

		for j, c := range line {
			if checkIfStar(string(c)) {
				specialCharMap[lineNo] = append(specialCharMap[lineNo], j)
			}
			lineSchema = append(lineSchema, string(c))
		}

		engineSchema = append(engineSchema, lineSchema)
		lineNo += 1
	}

	gearCount := map[string][]int{}
	for i, lineSchema := range engineSchema {
		numberStr := ""
		charAround := false
		gearLocation := ""

		for j, item := range lineSchema {
			_, err := strconv.Atoi(string(item))
			if err == nil {

				if !charAround {
					a, b := returnIfStarAround(engineSchema, specialCharMap, i, j)
					if a != -1 {
						charAround = true
						gearLocation = fmt.Sprintf("%d,%d", a, b)
					}
				}

				numberStr += item

				if j == len(lineSchema)-1 {
					number, err := strconv.Atoi(numberStr)
					if err != nil {
						fmt.Println("Error converting number", err)
					}

					if charAround {
						gearCount[gearLocation] = append(gearCount[gearLocation], number)
						gearLocation = ""
						charAround = false
					}
				}

			} else if len(numberStr) > 0 {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Println("Error converting number", err)
				}

				if charAround {
					gearCount[gearLocation] = append(gearCount[gearLocation], number)
					gearLocation = ""
					charAround = false
				}

				numberStr = ""
			} else {
				numberStr = ""
			}
		}
	}

	for _, v := range gearCount {
		if len(v) == 2 {
			sum += (v[0] * v[1])
		}
	}

	fmt.Println("Answer for part two is", sum)
}
