package day1

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day1/input.txt")

	sum := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		re := regexp.MustCompile("^\\D*(\\d)(?:.*?(\\d))?\\D*$")
		match := re.FindStringSubmatch(line)

		if match != nil && len(match) == 3 {
			if match[2] == "" {
				match[2] = match[1]
			}
			numberString := match[1] + match[2]
			number, numberErr := strconv.Atoi(numberString)

			if numberErr != nil {
				fmt.Println("Error:", numberErr)
				return
			}

			sum += number
		}
	}

	fmt.Println("Day 1, Part 1:", sum)
}
