package day03

import "strconv"

var vectorsToCheck = [][]int{
	{0, 1},
	{0, -1},
	{-1, 1},
	{-1, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{1, 0},
}

func currentCharIsntUsed(char rune) bool {
	return char == '.'
}

func resetNumbers(numberString string, numberStringIndexes []int) (string, []int) {
	numberString = ""
	numberStringIndexes = []int{}
	return numberString, numberStringIndexes
}

func currentCharIsSymbol(char rune) bool {
	_, numberErr := strconv.Atoi(string(char))
	return numberErr != nil
}

func indexIsInBounds(lines []string, nextXIndex int, nextYIndex int) bool {
	return nextXIndex >= 0 && nextXIndex < len(lines) && nextYIndex >= 0 && nextYIndex < len(lines[0])
}

func endOfLine(line string, j int) bool {
	return j == len(line)-1
}
