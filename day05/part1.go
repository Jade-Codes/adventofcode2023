package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day05/input.txt")

	currentNumbers := []int{}
	currentCategories := []category{}

	for i, line := range lines {

		if i == 0 {
			line = strings.Replace(line, "seeds:", "", -1)
			curentNumberStrings := strings.Fields(line)
			for _, numberString := range curentNumberStrings {
				number, _ := strconv.Atoi(numberString)
				currentNumbers = append(currentNumbers, number)
			}
			continue
		}

		if i == len(lines)-1 || line == "" {
			currentNumbers = calculateNextNumbers(currentNumbers, currentCategories)
			currentCategories = []category{}
			continue
		}

		if strings.Contains(line, ":") {
			continue
		}

		currentNumbers := strings.Fields(line)

		destinationRangeStart, _ := strconv.Atoi(currentNumbers[0])
		sourceRangeStart, _ := strconv.Atoi(currentNumbers[1])
		rangeCount, _ := strconv.Atoi(currentNumbers[2])

		currentCategory := category{
			destinationRangeStart: destinationRangeStart,
			destinationRangeEnd:   destinationRangeStart + rangeCount - 1,
			sourceRangeStart:      sourceRangeStart,
			sourceRangeEnd:        sourceRangeStart + rangeCount - 1,
			rangeCount:            rangeCount,
		}

		currentCategories = append(currentCategories, currentCategory)
	}

	fmt.Println("Day 5, Part 1:", slices.Min(currentNumbers))
}

func calculateNextNumbers(currentNumbers []int, currentCategories []category) []int {
	nextNumbers := []int{}

	for _, number := range currentNumbers {
		numberAlreadyAdded := false
		for _, currentCategory := range currentCategories {
			if isNumberInSourceRange(number, currentCategory) {
				nextNumbers = append(nextNumbers, number+currentCategory.destinationRangeStart-currentCategory.sourceRangeStart)
				numberAlreadyAdded = true
				break
			}
		}
		if !numberAlreadyAdded {
			nextNumbers = append(nextNumbers, number)
		}
	}

	return nextNumbers
}

func isNumberInSourceRange(number int, currentcategory category) bool {
	return number >= currentcategory.sourceRangeStart && number <= currentcategory.sourceRangeEnd
}
