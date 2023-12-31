package day09

import (
	"fmt"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day09/input.txt")

	total := 0

	for _, line := range lines {
		stringArray := strings.Fields(line)
		intArray := utils.SliceAtoi(stringArray)
		intArray = reverse(intArray)
		previousIntArray := getPreviousReverseRow(intArray)

		total += getFirstMissingNumber(intArray, previousIntArray)
	}

	fmt.Println("Day 9, Part 2:", total)

}

func getFirstMissingNumber(intArray []int, previousIntArray []int) int {
	for i := 0; i < len(intArray)-1; i++ {
		if intArray[i] != intArray[i+1] {
			newIntArray := getPreviousReverseRow(previousIntArray)

			previousMissingInt := getFirstMissingNumber(previousIntArray, newIntArray)
			previousIntArray = append(previousIntArray, previousMissingInt)
			break
		}
	}

	return intArray[len(intArray)-1] - previousIntArray[len(previousIntArray)-1]
}

func getPreviousReverseRow(intArray []int) []int {
	newIntArray := []int{}

	for i := 0; i < len(intArray)-1; i++ {
		newInt := intArray[i] - intArray[i+1]
		newIntArray = append(newIntArray, newInt)
	}

	return newIntArray
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
