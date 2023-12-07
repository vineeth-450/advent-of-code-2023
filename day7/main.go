package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

var strength map[string]int

var cardStrength = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var cardStrengthJoker = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
	"J": 0,
}

func partOne() {
	strength = cardStrength
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	scanner := bufio.NewScanner(file)

	rankMap := map[string][]string{}
	handToBetMap := map[string]int{}
	handCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		handType := getHandType(split[0])

		rankMap[handType] = append(rankMap[handType], split[0])
		bet, _ := strconv.Atoi(split[1])

		handToBetMap[split[0]] = bet
		handCount += 1
	}

	sum, tempSum := 0, 0

	orderOfHandTypes := []string{"5_OF_A_KIND", "4_OF_A_KIND", "FULL_HOUSE", "3_OF_A_KIND", "2_PAIR", "1_PAIR", "HIGH_CARD"}

	for _, ht := range orderOfHandTypes {
		handTypeSet := rankMap[ht]
		sort.Sort(Hand(handTypeSet))
		tempSum, handCount = calculateHandSetSum(handToBetMap, handTypeSet, handCount)
		sum += tempSum
	}

	fmt.Println("Answer for part one is", sum)
}

func calculateHandSetSum(handToBetMap map[string]int, handSet []string, rankStart int) (int, int) {
	sum := 0
	for _, h := range handSet {
		if handToBetMap[h] < 1 {
			fmt.Println("***")
		}
		sum += (rankStart * handToBetMap[h])
		rankStart -= 1
	}

	return sum, rankStart
}

func getHandType(hand string) string {
	cardMap := map[rune]int{}
	for _, c := range hand {
		cardMap[c] += 1
	}

	countArr := []int{}
	for _, v := range cardMap {
		countArr = append(countArr, v)
	}

	if compareIntArr(countArr, []int{5}) {
		return "5_OF_A_KIND"
	}
	if compareIntArr(countArr, []int{1, 4}) {
		return "4_OF_A_KIND"
	}
	if compareIntArr(countArr, []int{2, 3}) {
		return "FULL_HOUSE"
	}
	if compareIntArr(countArr, []int{1, 1, 3}) {
		return "3_OF_A_KIND"
	}
	if compareIntArr(countArr, []int{1, 2, 2}) {
		return "2_PAIR"
	}
	if compareIntArr(countArr, []int{1, 1, 1, 2}) {
		return "1_PAIR"
	}
	if compareIntArr(countArr, []int{1, 1, 1, 1, 1}) {
		return "HIGH_CARD"
	}

	return ""
}

func compareIntArr(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)

	for i, n := range a {
		if n != b[i] {
			return false
		}
	}

	return true
}

type Hand []string

func (h Hand) Len() int {
	return len(h)
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hand) Less(i, j int) bool {
	for m := 0; m < 5; m++ {
		if strength[string(h[i][m])] > strength[string(h[j][m])] {
			return true
		}
		if strength[string(h[i][m])] != strength[string(h[j][m])] {
			return false
		}
	}

	return false
}

func partTwo() {
	strength = cardStrengthJoker
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	scanner := bufio.NewScanner(file)

	rankMap := map[string][]string{}
	handToBetMap := map[string]int{}
	handCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		handType := getHandTypeJokerAct(split[0])

		rankMap[handType] = append(rankMap[handType], split[0])
		bet, _ := strconv.Atoi(split[1])

		handToBetMap[split[0]] = bet
		handCount += 1
	}

	sum, tempSum := 0, 0

	orderOfHandTypes := []string{"5_OF_A_KIND", "4_OF_A_KIND", "FULL_HOUSE", "3_OF_A_KIND", "2_PAIR", "1_PAIR", "HIGH_CARD"}

	for _, ht := range orderOfHandTypes {
		handTypeSet := rankMap[ht]
		sort.Sort(Hand(handTypeSet))
		tempSum, handCount = calculateHandSetSum(handToBetMap, handTypeSet, handCount)
		sum += tempSum
	}

	fmt.Println("Answer for part two is", sum)
}

func getHandTypeJokerAct(hand string) string {
	jCount := 0

	cardMap := map[rune]int{}
	for _, c := range hand {
		cardMap[c] += 1
		if string(c) == "J" {
			jCount += 1
		}
	}

	countArr := []int{}
	for k, v := range cardMap {
		if jCount != 5 && string(k) == "J" {
			continue
		}
		countArr = append(countArr, v)
	}

	if jCount != 5 {
		sort.Ints(countArr)
		countArr[len(countArr)-1] += jCount
	}

	if compareIntArr(countArr, []int{5}) {
		return "5_OF_A_KIND"
	}
	if compareIntArr(countArr, []int{1, 4}) {
		return "4_OF_A_KIND"
	}
	if compareIntArr(countArr, []int{2, 3}) {
		return "FULL_HOUSE"
	}
	if compareIntArr(countArr, []int{1, 1, 3}) {
		return "3_OF_A_KIND"
	}
	if compareIntArr(countArr, []int{1, 2, 2}) {
		return "2_PAIR"
	}
	if compareIntArr(countArr, []int{1, 1, 1, 2}) {
		return "1_PAIR"
	}
	if compareIntArr(countArr, []int{1, 1, 1, 1, 1}) {
		return "HIGH_CARD"
	}

	return ""
}
