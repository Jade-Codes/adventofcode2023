package day12

import (
	"fmt"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day12/input.txt")
	countValue := 0

	for _, line := range lines {
		sequence, numbers := parse(line)
		countValue += count(sequence, numbers)
	}

	fmt.Println("Day 12, Part 1:", countValue)
}

func parse(line string) (string, []int) {
	options := strings.Split(line, " ")
	sequence := options[0]
	numbers := utils.SliceAtoi(strings.Split(options[1], ","))
	return sequence, numbers
}
