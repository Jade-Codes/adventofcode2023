package day04

import (
	"fmt"
	"math"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day04/input.txt")

	power := float64(0)

	for i, line := range lines {
		scratchCardName := fmt.Sprintf("Card %d:", i+1)
		scratchCard := strings.Replace(line, scratchCardName, "", -1)

		winningNumbers := strings.Split(scratchCard, "|")[0]
		playingNumbers := strings.Split(scratchCard, "|")[1]

		winningNumbersArray := strings.Fields(winningNumbers)
		playingNumbersArray := strings.Fields(playingNumbers)

		numbersWon := -1

		for _, number := range winningNumbersArray {
			for _, playingNumber := range playingNumbersArray {
				if number == playingNumber {
					numbersWon++
				}
			}
		}

		if numbersWon > -1 {
			power += math.Pow(float64(2), float64(numbersWon))
		}
	}

	fmt.Println("Day 4, Part 1:", power)
}
