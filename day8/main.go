package main

import (
	"bufio"
	"fmt"
	"os"
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

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	route := scanner.Text()

	routeMap := map[string][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var source, left, right string
		fmt.Sscanf(line, "%s = (%3s, %3s)", &source, &left, &right)

		routeMap[source] = []string{left, right}
	}

	currentLoc := "AAA"
	turnCount := 0
	rounteIndex := 0
	for currentLoc != "ZZZ" {
		if rounteIndex == len(route) {
			rounteIndex = 0
		}
		t := string(route[rounteIndex])
		rounteIndex += 1

		if string(t) == "R" {
			currentLoc = routeMap[currentLoc][1]
		} else if string(t) == "L" {
			currentLoc = routeMap[currentLoc][0]
		}
		turnCount += 1

		if currentLoc == "ZZZ" {
			break
		}
	}

	fmt.Println("Answer for part one is", turnCount)
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	route := scanner.Text()
	fmt.Println("route len", len(route))

	routeMap := map[string][]string{}
	sourceLocations := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var source, left, right string
		fmt.Sscanf(line, "%s = (%3s, %3s)", &source, &left, &right)

		routeMap[source] = []string{left, right}

		if string(source[2]) == "A" {
			sourceLocations = append(sourceLocations, source)
		}
	}

	fmt.Println("source locations", len(sourceLocations), sourceLocations)
	destinationCount := 0
	turnCount := 0
	rounteIndex := 0
	for destinationCount != len(sourceLocations) {
		destinationCount = 0
		if rounteIndex == len(route) {
			rounteIndex = 0
		}
		t := string(route[rounteIndex])
		rounteIndex += 1

		// fmt.Print("route index", rounteIndex)

		for i, l := range sourceLocations {
			if string(t) == "R" {
				sourceLocations[i] = routeMap[l][1]
			} else if string(t) == "L" {
				sourceLocations[i] = routeMap[l][0]
			}

			if string(sourceLocations[i][2]) == "Z" {
				// fmt.Println("dest found")
				// fmt.Println(sourceLocations[i])
				destinationCount += 1
				if destinationCount > 2 {
					fmt.Println("dest count", destinationCount, sourceLocations, turnCount)
				}
			}
		}

		turnCount += 1
		// fmt.Println("dest count", destinationCount)
	}

	fmt.Println("Answer for part two is", turnCount)
}
