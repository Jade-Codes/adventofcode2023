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
	currentPosition := position{-1, -1, 'N'}

	for y, line := range lines {
		startingPointFound := false
		for x, value := range line {
			if value == 'S' {
				startingPointFound = true
				currentPosition = position{x, y, 'N'}
				break
			}
		}

		if startingPointFound {
			break
		}
	}
	counter := 0

	direction := nextDirection(currentPosition, lines, 'N')
	for currentPosition.value != 'S' {
		currentPosition = move(lines, direction, currentPosition)
		counter++
		direction = nextDirection(currentPosition, lines, direction)

	}

	fmt.Println("Day 10, Part 1:", counter/2)
}

func move(lines []string, direction rune, currentPosition position) position {
	if direction == 'L' {
		return moveLeft(currentPosition.value, currentPosition, lines)
	} else if direction == 'R' {
		return moveRight(currentPosition.value, currentPosition, lines)
	} else if direction == 'U' {
		return moveUp(currentPosition.value, currentPosition, lines)
	} else if direction == 'D' {
		return moveDown(currentPosition.value, currentPosition, lines)
	}
	return position{currentPosition.x, currentPosition.y, currentPosition.value}
}

func nextDirection(currentPosition position, lines []string, lastDirection rune) rune {
	fmt.Println(string(lastDirection))
	if currentPosition.x != 0 && lastDirection != 'R' {
		moveLeft := lines[currentPosition.y][currentPosition.x-1]
		stringValue := string(moveLeft)
		fmt.Println(stringValue)
		if moveLeft == '-' || moveLeft == 'L' || moveLeft == 'F' || moveLeft == 'S' {
			return 'L'
		}
	}
	if currentPosition.x != len(lines[currentPosition.y])-1 && lastDirection != 'L' {
		moveRight := lines[currentPosition.y][currentPosition.x+1]
		stringValue := string(moveRight)
		fmt.Println(stringValue)
		if moveRight == '-' || moveRight == 'J' || moveRight == '7' || moveRight == 'S' {
			return 'R'
		}
	}
	if currentPosition.y != 0 && lastDirection != 'D' {
		moveUp := lines[currentPosition.y-1][currentPosition.x]
		if moveUp == '|' || moveUp == 'F' || moveUp == '7' || moveUp == 'S' {
			return 'U'
		}
	}
	if currentPosition.y != len(lines)-1 && lastDirection != 'U' {
		moveDown := lines[currentPosition.y+1][currentPosition.x]
		if moveDown == '|' || moveDown == 'L' || moveDown == 'J' || moveDown == 'S' {
			return 'D'
		}
	}
	return lastDirection
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
