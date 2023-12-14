package day15

import (
	"fmt"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day15/input.txt")

	totalValue := 0
	for _, line := range lines {
		codes := strings.Split(line, ",")
		for _, code := range codes {
			totalValue += calculateHash(code)
		}

	}

	fmt.Println("Day 15, Part 1:", totalValue)
}

func calculateHash(label string) int {
	hash := 0
	for _, value := range label {
		ascii := int(value)
		hash += ascii
		hash *= 17
		hash %= 256
	}
	return hash
}
