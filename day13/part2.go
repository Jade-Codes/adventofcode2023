package day13

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day13/input.txt")
	newPattern := [][]rune{}
	patterns := [][][]rune{}
	columns := []int{}
	rows := []int{}
	for i, line := range lines {

		if i == len(lines)-1 {
			newPattern = append(newPattern, []rune(line))
			patterns = append(patterns, newPattern)
			break
		}

		if line == "" {
			patterns = append(patterns, newPattern)
			newPattern = [][]rune{}
			continue
		}
		newPattern = append(newPattern, []rune(line))
	}

	for _, pattern := range patterns {
		newPattern := [][]rune{}
		for i, row := range pattern {
			found := false
			for j, column := range row {
				if column == '#' {
					newPattern = make([][]rune, len(pattern))
					for i, r := range pattern {
						newPattern[i] = make([]rune, len(r))
						copy(newPattern[i], r)
					}
					newPattern[i][j] = '.'
				} else if column == '.' {
					newPattern = make([][]rune, len(pattern))
					for i, r := range pattern {
						newPattern[i] = make([]rune, len(r))
						copy(newPattern[i], r)
					}
					newPattern[i][j] = '#'
				}

				oldReflectionColumn := findColumnMirrors(pattern)
				oldReflectionRow := findRowMirrors(pattern)

				column := findColumnMirrors(newPattern)
				row := findRowMirrors(newPattern)

				if len(oldReflectionColumn) > 0 && len(column) > 0 {
					column = removeFromSlice(column, oldReflectionColumn[0])
				}

				if len(oldReflectionRow) > 0 && len(row) > 0 {
					row = removeFromSlice(row, oldReflectionRow[0])
				}

				if len(column) > 0 {
					for _, columnValue := range column {
						columns = append(columns, columnValue+1)
					}
					found = true
					break
				}

				if len(row) > 0 {
					for _, rowValue := range row {
						rows = append(rows, rowValue+1)
					}
					found = true
					break
				}
			}

			if found {
				break
			}
		}
	}

	total := 0
	for _, row := range rows {
		total += row * 100
	}

	for _, column := range columns {
		total += column
	}

	fmt.Println("Day 13, Part 1:", total)
}
