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
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	product := 1
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	timeStr := strings.Split(line, ":")[1]
	timeStr = strings.TrimSpace(timeStr)
	times := strings.Split(timeStr, " ")
	timesNum := []int{}
	for _, t := range times {
		if t == "" {
			continue
		}
		time, _ := strconv.Atoi(string(t))
		timesNum = append(timesNum, time)
	}

	scanner.Scan()
	line = scanner.Text()
	distanceStr := strings.Split(line, ":")[1]
	distanceStr = strings.TrimSpace(distanceStr)
	distances := strings.Split(distanceStr, " ")
	distancesNum := []int{}
	for _, d := range distances {
		if d == "" {
			continue
		}
		distance, _ := strconv.Atoi(string(d))
		distancesNum = append(distancesNum, distance)
	}

	for i, t := range timesNum {
		beatCount := 0
		for j := 1; j < t; j++ {
			distanceAchieved := j * (t - j)
			if distanceAchieved > distancesNum[i] {
				beatCount += 1
			}
		}

		product = product * beatCount
	}

	fmt.Println("Answer for part one is", product)
}
