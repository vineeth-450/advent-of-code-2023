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

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(line, " ")
		nums := []int{}
		for _, n := range numsStr {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		mat := [][]int{}
		mat = append(mat, nums)
		i, j := 0, 0
		diffArr := []int{}
		allZeros := true

		for {
			if j == len(mat[i])-1 {
				i += 1

				mat = append(mat, diffArr)
				diffArr = []int{}
				if allZeros {
					sum += calculateNextSequence(mat)
					break
				}

				allZeros = true
				j = 0
			}

			diff := (mat[i][j+1] - mat[i][j])
			diffArr = append(diffArr, diff)

			if diff != 0 {
				allZeros = false
			}

			j += 1
		}
	}

	fmt.Println("Answer for part one is", sum)
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error resing input file", err)
		return
	}

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(line, " ")
		nums := []int{}
		for _, n := range numsStr {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		mat := [][]int{}
		mat = append(mat, nums)
		i, j := 0, 0
		diffArr := []int{}
		allZeros := true

		for {
			if j == len(mat[i])-1 {
				i += 1

				mat = append(mat, diffArr)
				diffArr = []int{}
				if allZeros {
					sum += calculatePreviousSequence(mat)
					break
				}

				allZeros = true
				j = 0
			}

			diff := (mat[i][j+1] - mat[i][j])
			diffArr = append(diffArr, diff)

			if diff != 0 {
				allZeros = false
			}

			j += 1
		}
	}

	fmt.Println("Answer for part two is", sum)
}

func calculateNextSequence(mat [][]int) int {
	sum := 0
	for _, a := range mat {
		sum += a[len(a)-1]
	}

	return sum
}

func calculatePreviousSequence(mat [][]int) int {
	sum := 0

	i := len(mat) - 1
	for i != 0 {
		sum = mat[i-1][0] - sum
		i -= 1
	}

	return sum
}
