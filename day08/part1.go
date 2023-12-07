package day08

import (
	"fmt"
	"regexp"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

type directions struct {
	left, right string
}

func Part1() {
	lines := utils.GetLines("day08/input.txt")

	directions := lines[0]

	lettersMap := getMap(lines[2:])

	index := 0
	found := false
	currentLocation := "AAA"
	steps := 0
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
		steps++
		index++
		if currentLocation == "ZZZ" {
			found = true
		}
	}

	fmt.Println("Day 8, Part 1:", steps)

}

func getMap(lines []string) map[string]directions {
	m := make(map[string]directions)

	for _, line := range lines {
		re := regexp.MustCompile(`([A-Z]{3})\s*=\s*\(([^,]+),\s*([^)]+)\)`)
		matches := re.FindAllStringSubmatch(line, -1)

		m[matches[0][1]] = directions{matches[0][2], matches[0][3]}
	}
	return m

}
