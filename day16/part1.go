package day16

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

type Position struct {
	x                int
	y                int
	currentDirection rune
}

type Visited struct {
	x int
	y int
}

const north = '^'
const south = 'v'
const east = '>'
const west = '<'

const forwardReflection = '/'
const backwardReflection = '\\'
const straightReflection = '|'
const sideReflection = '-'

func Part1() {
	lines := utils.GetLines("day16/input.txt")
	matrix := utils.GetMatrix(lines)
	current := Position{0, 0, east}
	current.currentDirection = getNewDirections(current, matrix[current.y][current.x])[0]
	matrix[current.y][current.x] = current.currentDirection

	visited := map[Visited]bool{}
	visited[Visited{current.x, current.y}] = true

	getLocations(matrix, []Position{current}, visited)

	fmt.Println("Day 16, Part 1:", len(visited))
}

func getLocations(matrix [][]rune, positions []Position, visited map[Visited]bool) {
	for _, position := range positions {
		current := nextLocation(position)
		if current.x < 0 || current.y < 0 || current.x >= len(matrix[0]) || current.y >= len(matrix) {
			continue
		}

		visited[Visited{current.x, current.y}] = true

		if matrix[current.y][current.x] == north && current.currentDirection == north {
			continue
		}
		if matrix[current.y][current.x] == south && current.currentDirection == south {
			continue
		}
		if matrix[current.y][current.x] == east && current.currentDirection == east {
			continue
		}
		if matrix[current.y][current.x] == west && current.currentDirection == west {
			continue
		}

		if matrix[current.y][current.x] == '.' {
			matrix[current.y][current.x] = current.currentDirection
		}

		newDirections := getNewDirections(current, matrix[current.y][current.x])

		if len(newDirections) > 0 {
			newPositions := []Position{}
			for _, direction := range newDirections {
				newPositions = append(newPositions, Position{current.x, current.y, direction})
			}
			getLocations(matrix, newPositions, visited)
		}

	}
}

func nextLocation(current Position) Position {
	switch current.currentDirection {
	case north:
		current.y--
	case south:
		current.y++
	case east:
		current.x++
	case west:
		current.x--
	}
	return current
}

func getNewDirections(current Position, reflection rune) []rune {
	switch reflection {
	case forwardReflection:
		switch current.currentDirection {
		case north:
			return []rune{east}
		case south:
			return []rune{west}
		case east:
			return []rune{north}
		case west:
			return []rune{south}
		}
	case backwardReflection:
		switch current.currentDirection {
		case north:
			return []rune{west}
		case south:
			return []rune{east}
		case east:
			return []rune{south}
		case west:
			return []rune{north}
		}
	case straightReflection:
		switch current.currentDirection {
		case north:
			return []rune{north}
		case south:
			return []rune{south}
		case east:
			return []rune{north, south}
		case west:
			return []rune{north, south}
		}
	case sideReflection:
		switch current.currentDirection {
		case north:
			return []rune{east, west}
		case south:
			return []rune{east, west}
		case east:
			return []rune{east}
		case west:
			return []rune{west}
		}
	}
	return []rune{current.currentDirection}
}
