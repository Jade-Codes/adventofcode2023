package day03

import (
	"fmt"
	"strconv"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day03/input.txt")

	sumOfNumbers := 0

	for i, line := range lines {
		numberStringIndexes := []int{}
		numberString := ""
		for j, char := range line {

			if char == '.' && numberString != "" {
				sumOfNumbers += calculateNumberToAdd(lines, line, i, numberStringIndexes, numberString)
				numberString = ""
				numberStringIndexes = []int{}
			}

			if char == '.' {
				continue
			}

			number, numberErr := strconv.Atoi(string(char))

			if numberErr != nil && numberString != "" {
				number, _ := strconv.Atoi(numberString)
				sumOfNumbers += number

				numberString = ""
				numberStringIndexes = []int{}

				continue
			}

			numberStringIndexes = append(numberStringIndexes, j)
			numberString += strconv.Itoa(number)

			if j == len(line)-1 {
				sumOfNumbers += calculateNumberToAdd(lines, line, i, numberStringIndexes, numberString)
				numberString = ""
				numberStringIndexes = []int{}
			}
		}
	}

	fmt.Println("Day 3, Part 1:", sumOfNumbers)
}

func calculateNumberToAdd(lines []string, line string, rowIndex int, currentIndexes []int, numberString string) int {
	for i, columnIndex := range currentIndexes {
		canAdd := false

		if i == 0 {
			canAdd = canBeAdded(lines, line, rowIndex, columnIndex, vectorsToCheck)
		} else if i == len(currentIndexes)-1 {
			canAdd = canBeAdded(lines, line, rowIndex, columnIndex, vectorsToCheck)
		} else if i > 0 && i < len(currentIndexes)-1 {
			canAdd = canBeAdded(lines, line, rowIndex, columnIndex, vectorsToCheck)
		}

		if canAdd {
			number, _ := strconv.Atoi(numberString)
			return number
		}
	}
	return 0
}

func canBeAdded(lines []string, line string, rowIndex int, columnIndex int, vectorsToCheck [][]int) bool {
	for _, vector := range vectorsToCheck {
		nextXIndex := rowIndex + vector[0]
		nextYIndex := columnIndex + vector[1]

		if nextXIndex < 0 || nextXIndex >= len(lines) {
			continue
		}
		if nextYIndex < 0 || nextYIndex >= len(line) {
			continue
		}

		currentValue := lines[nextXIndex][nextYIndex]
		currentValueString := string(currentValue)

		if currentValueString == "." {
			continue
		}

		_, currentNumberErr := strconv.Atoi(currentValueString)

		if currentNumberErr != nil {
			return true
		}
	}

	return false
}
