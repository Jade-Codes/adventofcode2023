package day10

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day10/input.txt")

	counter := 0
	for y, line := range lines {
		inn := false
		for x, value := range line {
			firstDirection := 'N'
			currentPosition := position{x, y, value}
			currentDirection := nextDirection(currentPosition, lines, firstDirection)

			if currentDirection == 'N' {
				value = '.'
			}

			if value == '|' || value == 'J' || value == 'L' {
				inn = !inn
			}

			if (value == '.') && inn {
				counter++
			} else {
				fmt.Print(string(value))
			}
		}
		fmt.Println()
	}

	fmt.Println("Day 10, Part 2:", counter)
}
