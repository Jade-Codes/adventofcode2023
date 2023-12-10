package day11

import (
	"fmt"
	"math"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

const expansion = 1000000

func Part2() {
	lines := utils.GetLines("day11/input.txt")

	rowIndexesToAdd := []int{}
	columnIndexesToAdd := []int{}

	space := [][]rune{}

	for y, line := range lines {
		rowToBeAdded := true
		for _, value := range line {

			if value != '.' {
				rowToBeAdded = false
			}
		}
		if rowToBeAdded {
			rowIndexesToAdd = append(rowIndexesToAdd, y)
		}
	}

	for x := 0; x < len(lines[0]); x++ {
		columnToBeAdded := true
		for y := 0; y < len(lines); y++ {
			if lines[y][x] != '.' {
				columnToBeAdded = false
			}
		}
		if columnToBeAdded {
			columnIndexesToAdd = append(columnIndexesToAdd, x)
		}
	}

	for y, line := range lines {
		if containsInt(rowIndexesToAdd, y) {
			length := len(line)

			newRow := []rune{}

			for i := 0; i < length; i++ {
				newRow = append(newRow, 'X')
			}

			space = append(space, newRow)
		} else {
			space = append(space, []rune(line))
		}
	}

	for y := 0; y < len(space); y++ {
		for x := 0; x < len(space[0]); x++ {
			index := 0
			if containsInt(columnIndexesToAdd, x) {

				for i := 0; i < len(space); i++ {
					space[y][x] = 'Y'
				}

				index++
			}

		}
	}

	galaxies := [][]float64{}

	for y, line := range space {
		for x, value := range line {
			if value == '#' {
				galaxies = append(galaxies, []float64{float64(x), float64(y)})
			}
		}
	}

	distances := getDistancesInSpace(galaxies, space)

	total := 0.0

	for _, distance := range distances {
		total += distance
	}

	fmt.Println("Day 11, Part 2:", int(total))

}

func getDistancesInSpace(galaxies [][]float64, space [][]rune) []float64 {
	distances := []float64{}

	for x := 1; x < len(galaxies); x++ {
		distance := getManhattanDistanceInSpace(position{galaxies[0][0], galaxies[0][1]}, position{galaxies[x][0], galaxies[x][1]}, space)
		distances = append(distances, distance)
	}

	if len(galaxies) > 1 {
		nextGalaxies := galaxies[1:]
		distances = append(distances, getDistancesInSpace(nextGalaxies, space)...)
	}

	return distances
}

func getManhattanDistanceInSpace(a position, b position, space [][]rune) float64 {
	distance := 0

	minY := int(math.Min(a.y, b.y))
	maxY := int(math.Max(a.y, b.y))

	minX := int(math.Min(a.x, b.x))
	maxX := int(math.Max(a.x, b.x))

	for x := minX; x < maxX; x++ {
		if space[minY][x] == 'Y' {
			distance += expansion
		} else {
			distance++
		}
	}

	for y := minY; y < maxY; y++ {
		if space[y][minX] == 'X' {
			distance += expansion
		} else {
			distance++
		}
	}

	return float64(distance)
}
