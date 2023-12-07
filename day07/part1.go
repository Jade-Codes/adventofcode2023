package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day07/input.txt")
	games := []game{}

	for _, line := range lines {

		hand := strings.Fields(line)[0]
		bid, _ := strconv.Atoi(strings.Fields(line)[1])
		letterRanks := []int{}
		letterRanks = append(letterRanks, orderingMap[hand[0:1]], orderingMap[hand[1:2]], orderingMap[hand[2:3]], orderingMap[hand[3:4]], orderingMap[hand[4:5]])

		if isFiveOfAKind(hand) {
			games = append(games, game{hand, bid, 7, letterRanks})
			continue
		}

		if isFourOfAKind(hand) {
			games = append(games, game{hand, bid, 6, letterRanks})
			continue
		}

		if isFullHouse(hand) {
			games = append(games, game{hand, bid, 5, letterRanks})
			continue
		}

		if isThreeOfAKind(hand) {
			games = append(games, game{hand, bid, 4, letterRanks})
			continue
		}

		if isTwoPair(hand) {
			games = append(games, game{hand, bid, 3, letterRanks})
			continue
		}

		if isPair(hand) {
			games = append(games, game{hand, bid, 2, letterRanks})
			continue
		}

		if isHighCard(hand) {
			games = append(games, game{hand, bid, 1, letterRanks})
		}
	}

	sortGames(games)

	totalWinnings := getWinnings(games)

	fmt.Println("Day 7, Part 1:", totalWinnings)

}

var orderingMap = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}
