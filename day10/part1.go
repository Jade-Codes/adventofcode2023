package day10

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

type position struct {
	x, y  int
	value rune
}

const up = 'U'
const down = 'D'
const left = 'L'
const right = 'R'

func Part1() {
	lines := utils.GetLines("day10/input.txt")

	counter := 0
	startingPointFound := false

	currentPosition := findStartingPoint(lines)
	direction := nextDirection(currentPosition, lines, 'N')

	for !startingPointFound {
		currentPosition = move(lines, direction, currentPosition)
		counter++
		direction = nextDirection(currentPosition, lines, direction)

		if currentPosition.value == 'S' {
			startingPointFound = true
			break
		}
	}

	farthestDistance := counter / 2
	fmt.Println("Day 10, Part 1:", farthestDistance)
}

func findStartingPoint(lines []string) position {
	for y, line := range lines {
		for x, value := range line {
			if value == 'S' {
				return position{x, y, 'S'}
			}
		}
	}
	return position{-1, -1, 'N'}
}

func nextDirection(currentPosition position, lines []string, lastDirection rune) rune {
	if canMoveLeft(lines, lastDirection, currentPosition) {
		return left
	}
	if canMoveRight(lines, lastDirection, currentPosition) {
		return right
	}
	if canMoveUp(lines, lastDirection, currentPosition) {
		return up
	}
	if canMoveDown(lines, lastDirection, currentPosition) {
		return down
	}
	return lastDirection
}

func move(lines []string, direction rune, currentPosition position) position {
	if direction == left {
		return moveLeft(currentPosition.value, currentPosition, lines)
	} else if direction == right {
		return moveRight(currentPosition.value, currentPosition, lines)
	} else if direction == up {
		return moveUp(currentPosition.value, currentPosition, lines)
	} else if direction == down {
		return moveDown(currentPosition.value, currentPosition, lines)
	}
	return position{currentPosition.x, currentPosition.y, currentPosition.value}
}

func moveDown(value rune, currentPosition position, lines []string) position {
	return position{currentPosition.x, currentPosition.y + 1, rune(lines[currentPosition.y+1][currentPosition.x])}
}

func moveUp(value rune, currentPosition position, lines []string) position {
	return position{currentPosition.x, currentPosition.y - 1, rune(lines[currentPosition.y-1][currentPosition.x])}
}

func moveRight(value rune, currentPosition position, lines []string) position {
	return position{currentPosition.x + 1, currentPosition.y, rune(lines[currentPosition.y][currentPosition.x+1])}
}

func moveLeft(value rune, currentPosition position, lines []string) position {
	return position{currentPosition.x - 1, currentPosition.y, rune(lines[currentPosition.y][currentPosition.x-1])}
}

func canMoveLeft(lines []string, lastDirection rune, currentPosition position) bool {
	if currentPosition.x != 0 && lastDirection != 'R' &&
		(currentPosition.value == 'S' || currentPosition.value == '7' || currentPosition.value == '-' || currentPosition.value == 'J') {
		leftPosition := lines[currentPosition.y][currentPosition.x-1]
		if leftPosition == '-' || leftPosition == 'L' || leftPosition == 'F' || leftPosition == 'S' {
			return true
		}
	}
	return false
}

func canMoveRight(lines []string, lastDirection rune, currentPosition position) bool {
	if currentPosition.x != len(lines[currentPosition.y])-1 && lastDirection != 'L' &&
		(currentPosition.value == 'S' || currentPosition.value == 'F' || currentPosition.value == '-' || currentPosition.value == 'L') {
		rightPosition := lines[currentPosition.y][currentPosition.x+1]
		if rightPosition == '-' || rightPosition == 'J' || rightPosition == '7' || rightPosition == 'S' {
			return true
		}
	}
	return false
}

func canMoveUp(lines []string, lastDirection rune, currentPosition position) bool {
	if currentPosition.y != 0 && lastDirection != 'D' &&
		(currentPosition.value == 'S' || currentPosition.value == 'J' || currentPosition.value == '|' || currentPosition.value == 'L') {
		upPosition := lines[currentPosition.y-1][currentPosition.x]
		if upPosition == '|' || upPosition == 'F' || upPosition == '7' || upPosition == 'S' {
			return true
		}
	}
	return false
}

func canMoveDown(lines []string, lastDirection rune, currentPosition position) bool {
	if currentPosition.y != len(lines)-1 && lastDirection != 'U' &&
		(currentPosition.value == 'S' || currentPosition.value == 'F' || currentPosition.value == '|' || currentPosition.value == '7') {
		moveDown := lines[currentPosition.y+1][currentPosition.x]
		if moveDown == '|' || moveDown == 'L' || moveDown == 'J' || moveDown == 'S' {
			return true
		}
	}
	return false
}
