package day13

import (
	"fmt"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
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
		column := findColumnMirrors(pattern)
		row := findRowMirrors(pattern)
		if len(column) > 0 && len(row) > 0 {
			for _, columnValue := range column {
				for _, rowValue := range row {
					rows = append(rows, rowValue+1)
					columns = append(columns, columnValue+1)
				}
			}
		} else if len(column) > 0 {
			for _, columnValue := range column {
				columns = append(columns, columnValue+1)
			}
		} else if len(row) > 0 {
			for _, rowValue := range row {
				rows = append(rows, rowValue+1)
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

func findColumnMirrors(pattern [][]rune) []int {
	mirrors := []int{}
	newMirrors := []int{}
	for i := 0; i < len(pattern[0])-1; i++ {
		mirrors = append(mirrors, i)
		newMirrors = append(newMirrors, i)
	}

	for _, patternLine := range pattern {

		for _, char := range mirrors {

			isPotentialMirror := true
			mirrorStartingPoint := char
			reflectionStartingPoint := char + 1

			for isPotentialMirror {
				if patternLine[mirrorStartingPoint] != patternLine[reflectionStartingPoint] {
					isPotentialMirror = false
					newMirrors = removeFromSlice(newMirrors, char)
					break
				}
				if mirrorStartingPoint == 0 || reflectionStartingPoint == len(patternLine)-1 {
					break
				}
				mirrorStartingPoint--
				reflectionStartingPoint++
			}
		}
	}

	return newMirrors
}

func findRowMirrors(pattern [][]rune) []int {
	mirrors := []int{}
	newMirrors := []int{}
	for i := 0; i < len(pattern)-1; i++ {
		mirrors = append(mirrors, i)
		newMirrors = append(newMirrors, i)
	}

	for i := 0; i < len(pattern[0]); i++ {
		for _, char := range mirrors {

			isPotentialMirror := true
			mirrorStartingPoint := char
			relectionStartingPoint := char + 1

			for isPotentialMirror {
				if pattern[mirrorStartingPoint][i] != pattern[relectionStartingPoint][i] {
					isPotentialMirror = false
					newMirrors = removeFromSlice(newMirrors, char)
				}
				if mirrorStartingPoint == 0 || relectionStartingPoint == len(pattern)-1 {
					isPotentialMirror = false
				}
				mirrorStartingPoint--
				relectionStartingPoint++
			}

		}
	}

	return newMirrors
}

func removeFromSlice(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
