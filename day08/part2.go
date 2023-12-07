package day08

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day08/input.txt")

	directions := lines[0]

	lettersMap := getMap(lines[2:])

	locations := getStartingLocations(lettersMap)

	minStepsArray := []int{}
	for _, currentLocation := range locations {
		steps := 0
		found := false
		index := 0
		for !found {
			if index >= len(directions) {
				index = 0
			}

			step := directions[index]

			if step == 'L' {
				currentLocation = lettersMap[currentLocation].left
			}
			if step == 'R' {
				currentLocation = lettersMap[currentLocation].right
			}

			if currentLocation[2] == 'Z' {
				found = true
			}

			steps++
			index++
		}

		minStepsArray = append(minStepsArray, steps)
	}

	lowestCommonMultiple := 1

	for i := 0; i < len(minStepsArray); i++ {
		lowestCommonMultiple = getLeastCommonMultiple(lowestCommonMultiple, minStepsArray[i])
	}

	fmt.Println("Day 8, Part 2:", lowestCommonMultiple)
}

func getStartingLocations(m map[string]directions) []string {
	var locations []string
	for k, _ := range m {
		if k[2] == 'A' {
			locations = append(locations, k)
		}
	}
	return locations
}

func getLeastCommonMultiple(a, b int) int {
	return a * b / getGreatestCommonDivisor(a, b)
}

func getGreatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return getGreatestCommonDivisor(b, a%b)
}
