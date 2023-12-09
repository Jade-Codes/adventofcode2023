package day11

import (
	"fmt"
	"math"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

type position struct {
	x, y float64
}

func Part1() {
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
			columnIndexesLength := len(columnIndexesToAdd)
			columnIndexesToAdd = append(columnIndexesToAdd, x+columnIndexesLength)
		}
	}

	for y, line := range lines {
		if containsInt(rowIndexesToAdd, y) {
			length := len(line)

			newRow := []rune{}

			for i := 0; i < length; i++ {
				newRow = append(newRow, '.')
			}

			space = append(space, newRow)
		}
		space = append(space, []rune(line))
	}

	for y := 0; y < len(space); y++ {
		for x := 0; x < len(space[0]); x++ {
			index := 0
			if containsInt(columnIndexesToAdd, x) {

				newRow := make([]rune, len(space[y])+1)
				copy(newRow[:x], space[y][:x])

				newRow[x] = '.'

				copy(newRow[x+1:], space[y][x:])

				space[y] = newRow

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

	distances := getDistances(galaxies)

	total := 0.0

	for _, distance := range distances {
		total += distance
	}

	fmt.Println("Day 11, Part 1:", int(total))

}

func containsInt(array []int, value int) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

func getDistances(galaxies [][]float64) []float64 {
	distances := []float64{}

	for x := 1; x < len(galaxies); x++ {
		distance := getManhattanDistance(position{galaxies[0][0], galaxies[0][1]}, position{galaxies[x][0], galaxies[x][1]})
		distances = append(distances, distance)
	}

	if len(galaxies) > 1 {
		nextGalaxies := galaxies[1:]
		distances = append(distances, getDistances(nextGalaxies)...)
	}

	return distances
}

func getManhattanDistance(a position, b position) float64 {
	return math.Abs(a.x-b.x) + math.Abs(a.y-b.y)
}
