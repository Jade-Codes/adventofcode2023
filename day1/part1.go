package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	filePath := "day1/day1.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() //

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
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
