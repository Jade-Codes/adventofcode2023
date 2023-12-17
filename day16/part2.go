package day16

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day16/input.txt")
	matrix := utils.GetMatrix(lines)

	largestEnergisedCount := 0

	for y, line := range matrix {
		current := Position{0, y, east}
		matrix = utils.GetMatrix(lines)
		energisedCount := getEnergisedCount(current, matrix)

		if energisedCount > largestEnergisedCount {
			largestEnergisedCount = energisedCount
		}

		current = Position{len(line) - 1, y, west}
		energisedCount = getEnergisedCount(current, matrix)

		if energisedCount > largestEnergisedCount {
			largestEnergisedCount = energisedCount
		}
	}

	for x, _ := range matrix[0] {
		current := Position{x, 0, south}
		matrix = utils.GetMatrix(lines)
		energisedCount := getEnergisedCount(current, matrix)

		if energisedCount > largestEnergisedCount {
			largestEnergisedCount = energisedCount
		}

		current = Position{x, len(matrix) - 1, north}
		energisedCount = getEnergisedCount(current, matrix)
		if energisedCount > largestEnergisedCount {
			largestEnergisedCount = energisedCount
		}
	}

	fmt.Println("Day 16, Part 2:", largestEnergisedCount)
}

func getEnergisedCount(current Position, matrix [][]rune) int {
	currentDirections := getNewDirections(current, matrix[current.y][current.x])
	positions := []Position{}
	for _, direction := range currentDirections {
		positions = append(positions, Position{current.x, current.y, direction})
	}

	visited := map[Visited]bool{}
	visited[Visited{current.x, current.y}] = true

	getLocations(matrix, positions, visited)

	return len(visited)
}
