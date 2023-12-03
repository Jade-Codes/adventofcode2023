package day03

import (
	"fmt"
	"strconv"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day03/input.txt")

	sumOfNumbers := 0

	for i, line := range lines {
		for c, char := range line {

			if currentCharIsntUsed(char) {
				continue
			}

			if currentCharIsSymbol(char) {
				numbers := findSurroundingNumbers(lines, line, c, i)
				sumOfNumbers += getNumbersMultiplied(numbers)
			}

		}

	}

	fmt.Println("Day 3, Part 2:", sumOfNumbers)
}

func findSurroundingNumbers(lines []string, line string, symbolIndex int, rowIndex int) []int {
	numbers := []int{}

	checkedIndexesMatrix := [][]int{}

	for i := 0; i < len(lines); i++ {
		checkedIndexesMatrix = append(checkedIndexesMatrix, []int{})
	}

	for _, vector := range vectorsToCheck {

		nextXIndex := rowIndex + vector[0]
		nextYIndex := symbolIndex + vector[1]

		if !indexIsInBounds(lines, nextXIndex, nextYIndex) {
			continue
		}

		if isAlreadyChecked(checkedIndexesMatrix, nextXIndex, nextYIndex) {
			continue
		}

		nextNumberString := string(lines[nextXIndex][nextYIndex])
		nextNumber, nextNumberErr := strconv.Atoi(nextNumberString)

		if nextNumberErr != nil {
			continue
		}

		number, leftIndex, rightIndex := findFullNumberString(lines[nextXIndex], nextNumber, nextYIndex)

		checkedIndexesMatrix = appendToIndexesMatrix(checkedIndexesMatrix, nextXIndex, leftIndex, rightIndex)

		if number != 0 {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func findFullNumberString(line string, currentNumber int, currentIndex int) (int, int, int) {
	currentNumberString := strconv.Itoa(currentNumber)

	leftIndex, currentNumberString := getLeftIndex(line, currentIndex, currentNumberString)
	rightIndex, currentNumberString := getRightIndex(line, currentIndex, currentNumberString)

	currentNumber, currentNumberErr := strconv.Atoi(currentNumberString)

	if currentNumberErr != nil {
		return 0, leftIndex, rightIndex
	}

	return currentNumber, leftIndex, rightIndex
}

func appendToIndexesMatrix(checkedIndexesMatrix [][]int, xIndex int, leftIndex int, rightIndex int) [][]int {
	for i := leftIndex; i <= rightIndex; i++ {
		checkedIndexesMatrix[xIndex] = append(checkedIndexesMatrix[xIndex], i)
	}
	return checkedIndexesMatrix
}

func getNumbersMultiplied(numbers []int) int {
	if len(numbers) == 0 || len(numbers) == 1 {
		return 0
	}

	numbersMultiplied := 1
	for _, number := range numbers {
		numbersMultiplied *= number
	}

	return numbersMultiplied
}

func isAlreadyChecked(checkedIndexesMatrix [][]int, nextXIndex int, nextYIndex int) bool {
	for i, checkedIndexes := range checkedIndexesMatrix {
		if i == nextXIndex {
			for _, checkedIndex := range checkedIndexes {
				if checkedIndex == nextYIndex {
					return true
				}
			}
		}
	}
	return false
}

func getLeftIndex(line string, currentIndex int, currentNumberString string) (int, string) {
	leftEndNotFound := true
	leftIndex := int(currentIndex)

	for leftEndNotFound == true {
		if leftIndex == 0 {
			leftEndNotFound = false
			break
		}

		leftIndex--
		value := string(line[leftIndex])

		_, numberErr := strconv.Atoi(value)
		if numberErr != nil {
			leftEndNotFound = false
			break
		}

		currentNumberString = value + currentNumberString
	}

	return leftIndex, currentNumberString
}

func getRightIndex(line string, currentIndex int, currentNumberString string) (int, string) {
	rightEndNotFound := true
	rightIndex := int(currentIndex)

	for rightEndNotFound == true {
		if rightIndex == len(line)-1 {
			rightEndNotFound = false
			break
		}

		rightIndex++
		value := string(line[rightIndex])

		_, numberErr := strconv.Atoi(value)
		if numberErr != nil {
			rightEndNotFound = false
			break
		}

		currentNumberString = currentNumberString + value
	}

	return rightIndex, currentNumberString
}
