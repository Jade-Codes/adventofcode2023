package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

type Location struct {
	x, y int
}

func Part1() {
	lines := utils.GetLines("day18/input.txt")

	locations := []Location{}

	current := Location{0, 0}
	locations = append(locations, Location{current.x, current.y})

	for _, line := range lines {
		options := strings.Split(line, " ")

		direction := options[0]
		moveCount, _ := strconv.Atoi(options[1])

		for i := 0; i < moveCount; i++ {
			if direction == "R" {
				current.x++
			}
			if direction == "L" {
				current.x--
			}
			if direction == "U" {
				current.y++
			}
			if direction == "D" {
				current.y--
			}
			locations = append(locations, Location{current.x, current.y})
		}
	}

	shoelace := 0

	for i, location := range locations {

		if i == len(locations)-1 {
			shoelace += (location.x + locations[0].x) * (location.y - locations[0].y)
		} else {
			shoelace += (location.x + locations[i+1].x) * (location.y - locations[i+1].y)
		}
	}

	shoeLaceArea := shoelace / 2

	innerArea := shoeLaceArea - len(locations)/2 + 1

	totalArea := len(locations) + innerArea - 1

	fmt.Println("Day 18, Part 1:", totalArea)
}
