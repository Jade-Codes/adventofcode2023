package day18

import (
	"fmt"
	"math"
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

		if direction == "R" {
			current.x = current.x + moveCount
		}
		if direction == "L" {
			current.x = current.x - moveCount
		}
		if direction == "U" {
			current.y = current.y - moveCount
		}
		if direction == "D" {
			current.y = current.y + moveCount
		}
		locations = append(locations, Location{current.x, current.y})
	}

	shoelace := 0
	circumference := float64(0)

	for i, location := range locations {

		if i == len(locations)-1 {
			shoelace += (location.x - locations[0].x) * (location.y + locations[0].y)
			circumference += math.Abs(float64(location.x)-float64(locations[0].x)) + math.Abs(float64(location.y)-float64(locations[0].y))
		} else {
			shoelace += (location.x - locations[i+1].x) * (location.y + locations[i+1].y)
			circumference += math.Abs(float64(location.x)-float64(locations[i+1].x)) + math.Abs(float64(location.y)-float64(locations[i+1].y))
		}
	}

	shoeLaceArea := shoelace / 2

	innerArea := shoeLaceArea - int(circumference)/2 + 1

	totalArea := innerArea + int(circumference)

	fmt.Println("Day 18, Part 1:", totalArea)
}
