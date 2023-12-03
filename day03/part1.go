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

			if currentCharIsntUsed(char) && numberString != "" {
				sumOfNumbers += calculateNumberToAdd(lines, line, i, numberStringIndexes, numberString)
				numberString, numberStringIndexes = resetNumbers(numberString, numberStringIndexes)
			}

			if currentCharIsntUsed(char) {
				continue
			}

			if currentCharIsSymbol(char) {
				number, _ := strconv.Atoi(numberString)
				sumOfNumbers += number
				numberString, numberStringIndexes = resetNumbers(numberString, numberStringIndexes)

				continue
			}

			number, _ := strconv.Atoi(string(char))
			numberStringIndexes = append(numberStringIndexes, j)
			numberString += strconv.Itoa(number)

			if endOfLine(line, j) {
				sumOfNumbers += calculateNumberToAdd(lines, line, i, numberStringIndexes, numberString)

				numberString, numberStringIndexes = resetNumbers(numberString, numberStringIndexes)
			}
		}
	}

	fmt.Println("Day 3, Part 1:", sumOfNumbers)
}

func calculateNumberToAdd(lines []string, line string, rowIndex int, currentIndexes []int, numberString string) int {
	for _, columnIndex := range currentIndexes {
		if canBeAdded(lines, line, rowIndex, columnIndex, vectorsToCheck) {
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

		if !indexIsInBounds(lines, nextXIndex, nextYIndex) {
			continue
		}

		currentValue := lines[nextXIndex][nextYIndex]

		if currentCharIsntUsed(rune(currentValue)) {
			continue
		}

		if currentCharIsSymbol(rune(currentValue)) {
			return true
		}
	}

	return false
}
