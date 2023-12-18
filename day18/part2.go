package day18

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day18/input.txt")

	locations := []Location{}

	current := Location{0, 0}
	locations = append(locations, Location{current.x, current.y})

	for _, line := range lines {
		options := strings.Split(line, " ")

		direction, moveCount := convertHexToInt(options[2])

		if direction == 'R' {
			current.x = current.x + moveCount
		}
		if direction == 'L' {
			current.x = current.x - moveCount
		}
		if direction == 'U' {
			current.y = current.y - moveCount
		}
		if direction == 'D' {
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

	fmt.Println("Day 18, Part 2:", totalArea)
}

func convertHexToInt(hex string) (rune, int) {

	hex = strings.Replace(hex, "(", "", -1)
	hex = strings.Replace(hex, ")", "", -1)

	var start = hex[1 : len(hex)-1]
	var end = hex[len(hex)-1:]

	steps, _ := strconv.ParseInt(start, 16, 64)
	direction, _ := strconv.ParseInt(end, 16, 64)

	if direction == 0 {
		return 'R', int(steps)
	}
	if direction == 1 {
		return 'D', int(steps)
	}
	if direction == 2 {
		return 'L', int(steps)
	}
	if direction == 3 {
		return 'U', int(steps)
	}

	return 'R', 0
}
