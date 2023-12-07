package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day07/input.txt")
	games := []game{}

	for _, line := range lines {

		hand := strings.Fields(line)[0]
		bid, _ := strconv.Atoi(strings.Fields(line)[1])

		commonString := getMostCommonStringElement(hand)
		currentHand := strings.Replace(hand, "J", commonString, -1)

		letterRanks := []int{}
		letterRanks = append(letterRanks, jokerOrderingMap[hand[0:1]], jokerOrderingMap[hand[1:2]], jokerOrderingMap[hand[2:3]], jokerOrderingMap[hand[3:4]], jokerOrderingMap[hand[4:5]])

		if isFiveOfAKind(currentHand) {
			games = append(games, game{hand, bid, 7, letterRanks})
			continue
		}

		if isFourOfAKind(currentHand) {
			games = append(games, game{hand, bid, 6, letterRanks})
			continue
		}

		if isFullHouse(currentHand) {
			games = append(games, game{hand, bid, 5, letterRanks})
			continue
		}

		if isThreeOfAKind(currentHand) {
			games = append(games, game{hand, bid, 4, letterRanks})
			continue
		}

		if isTwoPair(currentHand) {
			games = append(games, game{hand, bid, 3, letterRanks})
			continue
		}

		if isPair(currentHand) {
			games = append(games, game{hand, bid, 2, letterRanks})
			continue
		}

		if isHighCard(currentHand) {
			games = append(games, game{hand, bid, 1, letterRanks})
		}
	}

	sortGames(games)

	totalWinnings := getWinnings(games)

	fmt.Println("Day 7, Part 2:", totalWinnings)

}

var jokerOrderingMap = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func getMostCommonStringElement(input string) string {
	strings := strings.Split(input, "")
	counts := make(map[string]int)

	for _, str := range strings {
		counts[str]++
	}

	var mostCommonValue string
	var maxCount int

	for value, count := range counts {
		if count > maxCount {
			if value != "J" {
				mostCommonValue = value
				maxCount = count
			}
		}
	}

	return mostCommonValue
}
