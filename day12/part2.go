package day12

import (
	"fmt"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day12/input.txt")
	countValue := 0

	for _, line := range lines {
		sequence, numbers := parseFold(line)
		countValue += count(sequence, numbers)
	}

	fmt.Println("Day 12, Part 2:", countValue)
}

func parseFold(line string) (string, []int) {

	options := strings.Split(line, " ")

	sequence := options[0]
	numbersString := options[1]

	for i := 0; i < 4; i++ {
		sequence = sequence + "?" + options[0]
	}

	for i := 0; i < 4; i++ {
		numbersString = numbersString + "," + options[1]
	}

	numbers := utils.SliceAtoi(strings.Split(numbersString, ","))
	return sequence, numbers
}
