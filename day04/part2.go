package day04

import (
	"fmt"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day04/input.txt")

	dictionary := map[int]float64{}

	for i, line := range lines {
		currentCard := i + 1

		scratchCardName := fmt.Sprintf("Card %d:", currentCard)
		scratchCard := strings.Replace(line, scratchCardName, "", -1)

		winningNumbers := strings.Split(scratchCard, "|")[0]
		playingNumbers := strings.Split(scratchCard, "|")[1]

		winningNumbersArray := strings.Fields(winningNumbers)
		playingNumbersArray := strings.Fields(playingNumbers)

		numbersWon := float64(0)

		dictionary[currentCard] += float64(1)

		for _, number := range winningNumbersArray {
			for _, playingNumber := range playingNumbersArray {

				if number == playingNumber {
					numbersWon++
				}
			}

		}

		if numbersWon > 0 {
			currentValue := dictionary[currentCard]
			nextCard := currentCard + 1

			for j := nextCard; j < nextCard+int(numbersWon); j++ {
				nextValue := dictionary[j]

				dictionary[j] = nextValue + currentValue
			}
		}
	}

	sumOfWinnings := 0
	for _, value := range dictionary {
		sumOfWinnings += int(value)
	}

	fmt.Println("Day 4, Part 2:", sumOfWinnings)
}
