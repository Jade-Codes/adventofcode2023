package day09

import (
	"fmt"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day09/input.txt")

	total := 0

	for _, line := range lines {
		stringArray := strings.Fields(line)
		intArray := utils.SliceAtoi(stringArray)
		previousIntArray := getPreviousRow(intArray)

		total += getLastMissingNumber(intArray, previousIntArray)
	}

	fmt.Println("Day 9, Part 1:", total)
}

func getLastMissingNumber(intArray []int, previousIntArray []int) int {
	for i := 0; i < len(intArray)-1; i++ {
		if intArray[i] != intArray[i+1] {
			newIntArray := getPreviousRow(previousIntArray)

			newInt := getLastMissingNumber(previousIntArray, newIntArray)
			previousIntArray = append(previousIntArray, newInt)
			break
		}
	}

	return intArray[len(intArray)-1] + previousIntArray[len(previousIntArray)-1]
}
func getPreviousRow(intArray []int) []int {
	newIntArray := []int{}

	for i := 0; i < len(intArray)-1; i++ {
		newInt := intArray[i+1] - intArray[i]
		newIntArray = append(newIntArray, newInt)
	}

	return newIntArray
}
