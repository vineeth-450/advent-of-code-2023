package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOne()
}

type rule struct {
	leftOperand  string
	operator     string
	rightOperand string
	destination  string
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	scanner := bufio.NewScanner(file)

	workflows := make(map[string][]rule)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		split := strings.Split(line, "{")
		wfName := split[0]
		wfRules := strings.TrimSuffix(split[1], "}")

		for _, rl := range strings.Split(wfRules, ",") {
			s := strings.Split(rl, ":")
			var newRule rule
			if len(s) > 1 {
				dest := s[1]
				re := regexp.MustCompile("[<>]")

				o := re.Split(s[0], -1)

				if len(o) > 1 {
					newRule = rule{o[0], re.FindString(s[0]), o[1], dest}
				}
			} else {
				newRule = rule{"", "", "", s[0]}
			}

			workflows[wfName] = append(workflows[wfName], newRule)
		}

	}

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := make(map[string]int)
		valueSum := 0
		for _, part := range strings.Split(strings.Trim(line, "{}"), ",") {
			s := strings.Split(part, "=")
			value, _ := strconv.Atoi(s[1])
			parts[s[0]] = value
			valueSum += value
		}

		if getAcceptance(workflows, parts, "in") {
			sum += valueSum
		}
	}

	fmt.Println("Answer for part one is", sum)
}

func getAcceptance(wfs map[string][]rule, parts map[string]int, wfName string) bool {
	if wfName == "R" {
		return false
	} else if wfName == "A" {
		return true
	}

	rules := wfs[wfName]
	for _, r := range rules {
		if r.operator == "" {
			return getAcceptance(wfs, parts, r.destination)
		}

		rightOperantInt, _ := strconv.Atoi(r.rightOperand)
		if operationMap[r.operator](parts[r.leftOperand], rightOperantInt) {
			return getAcceptance(wfs, parts, r.destination)
		}
	}

	return false
}

var operationMap = map[string]func(x, y int) bool{
	"<": lesserThan,
	">": greaterThan,
}

func lesserThan(x, y int) bool {
	return x < y
}

func greaterThan(x, y int) bool {
	return x > y
}
