package day04

import (
	"fmt"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day04/input.txt")

	cardDictionary := map[int]int{}

	for i, line := range lines {
		currentCard := i + 1

		scratchCardName := fmt.Sprintf("Card %d:", currentCard)
		scratchCard := strings.Replace(line, scratchCardName, "", -1)

		winningNumbers := strings.Split(scratchCard, "|")[0]
		playingNumbers := strings.Split(scratchCard, "|")[1]

		winningNumbersArray := strings.Fields(winningNumbers)
		playingNumbersArray := strings.Fields(playingNumbers)

		numbersWon := getNumbersWon(winningNumbersArray, playingNumbersArray)

		cardDictionary[currentCard] += 1
		cardDictionary = getNextCard(currentCard, numbersWon, cardDictionary)
	}

	sumOfCards := getSumOfCards(cardDictionary)

	fmt.Println("Day 4, Part 2:", sumOfCards)
}

func getNextCard(currentCard int, numbersWon int, cardDictionary map[int]int) map[int]int {
	if numbersWon > 0 {
		currentAmount := cardDictionary[currentCard]
		nextCard := currentCard + 1

		for j := nextCard; j < nextCard+int(numbersWon); j++ {
			nextAmount := cardDictionary[j]

			cardDictionary[j] = nextAmount + currentAmount
		}
	}
	return cardDictionary
}

func getSumOfCards(cardDictionary map[int]int) int {
	sumOfCards := 0
	for _, card := range cardDictionary {
		sumOfCards += card
	}
	return sumOfCards
}
