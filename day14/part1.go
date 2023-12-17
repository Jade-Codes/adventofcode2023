package day14

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day14/input.txt")

	matrix := utils.GetMatrix(lines)

	for y, line := range matrix {
		for x, _ := range line {
			matrix = move(matrix, x, y)
		}
	}

	totalLoad := 0
	for y, line := range matrix {
		currentRow := len(matrix) - y
		for x, _ := range line {
			if matrix[y][x] == 'O' {
				totalLoad += currentRow
			}
		}
	}

	fmt.Println("Day 14, Part 1:", totalLoad)
}

func move(matrix [][]rune, x int, y int) [][]rune {
	if matrix[y][x] == 'O' {
		canMove := true
		for canMove {

			if y == 0 {
				canMove = false
				break
			}

			y--

			if matrix[y][x] == '.' {
				matrix[y][x] = 'O'
				matrix[y+1][x] = '.'
			}

			if matrix[y][x] == '#' {
				canMove = false
				break
			}
		}
	}
	return matrix
}
