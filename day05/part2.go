package day05

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day05/input.txt")

	currentNumbers := [][]int{}
	currentCategories := []category{}

	for i, line := range lines {

		if i == 0 {
			currentNumbers = getSeedNumbers(currentNumbers, line)
			continue
		}

		if strings.Contains(line, ":") {
			continue
		}

		if isNextStep(lines, line, i) {
			currentNumbers = getCurrentNumbers(currentNumbers, currentCategories)
			currentCategories = []category{}
			continue
		}

		currentCategories = addToCategories(line, currentCategories)
	}

	minNumber := getMinNumber(currentNumbers)

	fmt.Println("Day 5, Part 2:", minNumber)
}

func addMissingNumbers(currentNumbers [][]int, start int, end int) [][]int {
	if start < end {
		currentNumbers = append(currentNumbers, []int{start, end})
	}
	return currentNumbers
}

func getCurrentNumbers(currentNumbers [][]int, currentCategories []category) [][]int {
	tempNumbers := [][]int{}
	for len(currentNumbers) > 0 {
		mapped := false
		currentNumber := currentNumbers[0]
		sourceStart := currentNumber[0]
		sourceEnd := currentNumber[1]
		for _, currentCategory := range currentCategories {
			destStart := int(math.Max(float64(currentCategory.sourceRangeStart), float64(sourceStart)))
			destEnd := int(math.Min(float64(currentCategory.sourceRangeEnd), float64(sourceEnd)))

			if destStart < destEnd {
				newStart := destStart + currentCategory.destinationRangeStart - currentCategory.sourceRangeStart
				newEnd := destEnd + currentCategory.destinationRangeStart - currentCategory.sourceRangeStart
				tempNumbers = append(tempNumbers, []int{newStart, newEnd})

				currentNumbers = addMissingNumbers(currentNumbers, destEnd, sourceEnd)
				currentNumbers = addMissingNumbers(currentNumbers, sourceStart, destStart)

				mapped = true
				break
			}
		}

		if !mapped {
			tempNumbers = append(tempNumbers, currentNumber)
		}

		currentNumbers = currentNumbers[1:]
	}

	return tempNumbers
}

func getMinNumber(currentNumbers [][]int) int {
	minNumber := currentNumbers[0][0]
	for _, currentNumber := range currentNumbers {
		if currentNumber[0] < minNumber {
			minNumber = currentNumber[0]
		}
	}
	return minNumber
}

func getSeedNumbers(seedNumbers [][]int, line string) [][]int {
	line = strings.Replace(line, "seeds:", "", -1)
	currentNumberStrings := strings.Fields(line)

	for j := 0; j < len(currentNumberStrings); j += 2 {
		number, _ := strconv.Atoi(currentNumberStrings[j])
		numberRange, _ := strconv.Atoi(currentNumberStrings[j+1])
		seedNumbers = append(seedNumbers, []int{number, number + numberRange})
	}

	return seedNumbers
}

func isNextStep(lines []string, line string, index int) bool {
	if index == len(lines)-1 || line == "" {
		return true
	}
	return false
}
