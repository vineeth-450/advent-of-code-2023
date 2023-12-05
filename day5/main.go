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

type sourceDestinationMappings []map[string]int

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	seedStr := strings.Split(line, ":")[1]
	seedStr = strings.TrimSpace(seedStr)
	seeds := strings.Split(seedStr, " ")

	seedNumbers := []int{}
	for _, s := range seeds {
		seedNumber, _ := strconv.Atoi(string(s))
		seedNumbers = append(seedNumbers, seedNumber)
	}

	mappings := make(map[string]sourceDestinationMappings)

	currentMapping := ""

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Split(line, " ")[1] == "map:" {
			currentMapping = strings.Split(line, " ")[0]
			continue
		}

		mappingStr := strings.Split(line, " ")
		destinationStart, _ := strconv.Atoi(string(mappingStr[0]))
		sourceStart, _ := strconv.Atoi(string(mappingStr[1]))
		rangeLength, _ := strconv.Atoi(string(mappingStr[2]))
		mapping := map[string]int{
			"destinationStart": destinationStart,
			"sourceStart":      sourceStart,
			"rangeLength":      rangeLength,
		}

		mappings[currentMapping] = append(mappings[currentMapping], mapping)
	}

	seedNumbers = mapToDestination(mappings["seed-to-soil"], seedNumbers)
	seedNumbers = mapToDestination(mappings["soil-to-fertilizer"], seedNumbers)
	seedNumbers = mapToDestination(mappings["fertilizer-to-water"], seedNumbers)
	seedNumbers = mapToDestination(mappings["water-to-light"], seedNumbers)
	seedNumbers = mapToDestination(mappings["light-to-temperature"], seedNumbers)
	seedNumbers = mapToDestination(mappings["temperature-to-humidity"], seedNumbers)
	seedNumbers = mapToDestination(mappings["humidity-to-location"], seedNumbers)

	min := seedNumbers[0]
	for i := 1; i < len(seedNumbers); i++ {
		if seedNumbers[i] < min {
			min = seedNumbers[i]
		}
	}

	fmt.Println("Answer for part one is", min)
}

func mapToDestination(mappings sourceDestinationMappings, seedNumbers []int) []int {
	newSeedNumbers := []int{}

	for _, n := range seedNumbers {
		mappingFound := false
		// fmt.Println("Processing", n)
		for _, m := range mappings {
			// fmt.Println("source", m["sourceStart"], "range", m["rangeLength"])
			if n >= m["sourceStart"] && n < m["sourceStart"]+m["rangeLength"] {
				// fmt.Println(n, " in range", m["sourceStart"], m["rangeLength"])
				newSeedNumbers = append(newSeedNumbers, m["destinationStart"]+(n-m["sourceStart"]))
				mappingFound = true
				break
			}
		}

		if !mappingFound {
			newSeedNumbers = append(newSeedNumbers, n)
		}

		mappingFound = false
	}

	return newSeedNumbers
}
