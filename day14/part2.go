package day14

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day14/input.txt")

	matrix := getMatrix(lines)
	matrixes := [][][]rune{}

	cycles := 1000000000

	for i := 0; i < cycles; i++ {
		matrix = cycle(matrix)

		if i%1000 == 0 {
			if matrixesContain(matrixes, matrix) {
				pattern, start := findRepeatablePattern(matrixes)
				cycles = (cycles - start) % pattern
				matrix = matrixes[start+cycles]
				break
			}
		}

		newMatrix := make([][]rune, len(matrix))
		for i := range matrix {
			newMatrix[i] = make([]rune, len(matrix[i]))
			copy(newMatrix[i], matrix[i])
		}
		matrixes = append(matrixes, newMatrix)
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

	fmt.Println("Day 14, Part 2:", totalLoad)
}

func cycle(matrix [][]rune) [][]rune {

	newMatrix := make([][]rune, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]rune, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}

	for i := 0; i < 4; i++ {
		up := i%4 == 0
		left := i%4 == 1
		down := i%4 == 2
		right := i%4 == 3

		if up {

			for y, line := range matrix {
				for x, _ := range line {
					matrix = moveUp(matrix, x, y)
				}
			}

		}

		if left {

			for y, line := range matrix {
				for x, _ := range line {
					matrix = moveLeft(matrix, x, y)
				}
			}
		}

		if down {
			for y := len(matrix) - 1; y >= 0; y-- {
				for x := len(matrix[y]) - 1; x >= 0; x-- {
					matrix = moveDown(matrix, x, y)
				}
			}
		}

		if right {
			for y := len(matrix) - 1; y >= 0; y-- {
				for x := len(matrix[y]) - 1; x >= 0; x-- {
					matrix = moveRight(matrix, x, y)
				}
			}
		}
	}

	return matrix
}

func canMoveLeft(matrix [][]rune, x int, y int) bool {
	if x == 0 {
		return false
	}

	if matrix[y][x-1] == '.' {
		return true
	}

	if matrix[y][x-1] == '#' {
		return false
	}

	if matrix[y][x-1] == 'O' {
		return canMoveLeft(matrix, x-1, y)
	}

	return false
}

func canMoveRight(matrix [][]rune, x int, y int) bool {
	if x == len(matrix[y])-1 {
		return false
	}

	if matrix[y][x+1] == '.' {
		return true
	}

	if matrix[y][x+1] == '#' {
		return false
	}

	if matrix[y][x+1] == 'O' {
		return canMoveRight(matrix, x+1, y)
	}

	return false
}

func canMoveUp(matrix [][]rune, x int, y int) bool {
	if y == 0 {
		return false
	}

	if matrix[y-1][x] == '.' {
		return true
	}

	if matrix[y-1][x] == '#' {
		return false
	}

	if matrix[y-1][x] == 'O' {
		return canMoveUp(matrix, x, y-1)
	}

	return false
}

func canMoveDown(matrix [][]rune, x int, y int) bool {
	if y == len(matrix)-1 {
		return false
	}

	if matrix[y+1][x] == '.' {
		return true
	}

	if matrix[y+1][x] == '#' {
		return false
	}

	if matrix[y+1][x] == 'O' {
		return canMoveDown(matrix, x, y+1)
	}

	return false
}

func findNextAvailableDownPosition(matrix [][]rune, x int, y int) (int, int) {
	if y == len(matrix)-1 {
		return y, x
	}

	if matrix[y+1][x] == '.' {
		return y + 1, x
	}

	if matrix[y+1][x] == '#' {
		return y, x
	}

	if matrix[y+1][x] == 'O' {
		return findNextAvailableDownPosition(matrix, x, y+1)
	}

	return y, x
}

func findNextAvailableRightPosition(matrix [][]rune, x int, y int) (int, int) {
	if x == len(matrix[y])-1 {
		return y, x
	}

	if matrix[y][x+1] == '.' {
		return y, x + 1
	}

	if matrix[y][x+1] == '#' {
		return y, x
	}

	if matrix[y][x+1] == 'O' {
		return findNextAvailableRightPosition(matrix, x+1, y)
	}

	return y, x
}

func moveUp(matrix [][]rune, x int, y int) [][]rune {
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

func moveLeft(matrix [][]rune, x int, y int) [][]rune {
	if matrix[y][x] == 'O' {
		canMove := true
		for canMove {

			if x == 0 {
				canMove = false
				break
			}

			x--

			if matrix[y][x] == '.' {
				matrix[y][x] = 'O'
				matrix[y][x+1] = '.'
			}

			if matrix[y][x] == '#' {
				canMove = false
				break
			}
		}
	}
	return matrix
}

func moveRight(matrix [][]rune, x int, y int) [][]rune {
	if matrix[y][x] == 'O' {
		canMove := true
		for canMove {

			if x == len(matrix[y])-1 {
				canMove = false
				break
			}

			x++

			if matrix[y][x] == '.' {
				matrix[y][x] = 'O'
				matrix[y][x-1] = '.'
			}

			if matrix[y][x] == '#' {
				canMove = false
				break
			}
		}
	}
	return matrix
}

func moveDown(matrix [][]rune, x int, y int) [][]rune {
	if matrix[y][x] == 'O' {
		canMove := true
		for canMove {

			if y == len(matrix)-1 {
				canMove = false
				break
			}

			y++

			if matrix[y][x] == '.' {
				matrix[y][x] = 'O'
				matrix[y-1][x] = '.'
			}

			if matrix[y][x] == '#' {
				canMove = false
				break
			}
		}
	}
	return matrix
}

func matrixesContain(matrixes [][][]rune, matrix [][]rune) bool {
	for _, m := range matrixes {
		if matrixMatch(m, matrix) {
			return true
		}
	}
	return false
}

func matrixMatch(matrix1 [][]rune, matrix2 [][]rune) bool {
	for y, line := range matrix1 {
		for x, _ := range line {
			if matrix1[y][x] != matrix2[y][x] {
				return false
			}
		}
	}
	return true
}

func findRepeatablePattern(matrixes [][][]rune) (int, int) {
	matrix := matrixes[len(matrixes)-1]

	for i := len(matrixes) - 2; i >= 0; i-- {
		if matrixMatch(matrixes[i], matrix) {
			return len(matrixes) - i, i
		}
	}

	return 0, 0
}
