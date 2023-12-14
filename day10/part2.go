package day10

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day10/input.txt")

	counter := 0
	startingPointFound := false

	currentPosition := findStartingPoint(lines)
	direction := nextDirection(currentPosition, lines, 'N')
	positions := []position{}
	positions = append(positions, currentPosition)

	for !startingPointFound {
		currentPosition = move(lines, direction, currentPosition)
		counter++
		direction = nextDirection(currentPosition, lines, direction)

		if currentPosition.value == 'S' {
			startingPointFound = true
			break
		}

		positions = append(positions, currentPosition)
	}

	shoelace := 0

	for i := 0; i < len(positions); i++ {

		if i == len(positions)-1 {
			shoelace += (positions[i].x + positions[0].x) * (positions[i].y - positions[0].y)
		} else {
			shoelace += (positions[i].x + positions[i+1].x) * (positions[i].y - positions[i+1].y)
		}
	}

	shoeLaceArea := shoelace / 2

	innerArea := shoeLaceArea - len(positions)/2 + 1

	fmt.Println("Day 10, Part 2:", innerArea)
}
