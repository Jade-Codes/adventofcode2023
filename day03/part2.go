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

			if string(char) == "." {
				continue
			}

			_, numberErr := strconv.Atoi(string(char))

			if numberErr != nil {
				numbers := findSurroundingNumbers(lines, line, c, i)
				if len(numbers) > 1 {
					numbersMultiplied := 1
					for _, number := range numbers {
						numbersMultiplied *= number
					}
					sumOfNumbers += numbersMultiplied
				}
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

		if nextXIndex < 0 || nextXIndex >= len(lines) {
			continue
		}

		if nextYIndex < 0 || nextYIndex >= len(line) {
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

		for i := leftIndex; i <= rightIndex; i++ {
			checkedIndexesMatrix[nextXIndex] = append(checkedIndexesMatrix[nextXIndex], i)
		}

		if number != 0 {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func findFullNumberString(line string, currentNumber int, currentIndex int) (int, int, int) {
	currentNumberString := strconv.Itoa(currentNumber)

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

	currentNumber, currentNumberErr := strconv.Atoi(currentNumberString)

	if currentNumberErr != nil {
		return 0, leftIndex, rightIndex
	}

	return currentNumber, leftIndex, rightIndex
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
