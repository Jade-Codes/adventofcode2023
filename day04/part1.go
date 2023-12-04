package day04

import (
	"fmt"
	"math"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day04/input.txt")

	totalPoints := 0

	for i, line := range lines {
		scratchCardName := fmt.Sprintf("Card %d:", i+1)
		scratchCard := strings.Replace(line, scratchCardName, "", -1)

		winningNumbers := strings.Split(scratchCard, "|")[0]
		playingNumbers := strings.Split(scratchCard, "|")[1]

		winningNumbersArray := strings.Fields(winningNumbers)
		playingNumbersArray := strings.Fields(playingNumbers)

		numbersWon := getNumbersWon(winningNumbersArray, playingNumbersArray)

		if numbersWon > 0 {
			totalPoints += getPoints(numbersWon)
		}
	}

	fmt.Println("Day 4, Part 1:", totalPoints)
}

func getPoints(numbersWon int) int {
	return int(math.Pow(float64(2), float64(numbersWon-1)))
}
