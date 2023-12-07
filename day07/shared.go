package day07

import (
	"sort"
	"strings"
)

type game struct {
	card        string
	bid, rank   int
	letterRanks []int
}

func isFiveOfAKind(hand string) bool {
	for i := 0; i < len(hand); i++ {
		if hand[i] != hand[0] {
			return false
		}
	}
	return true
}

func isFourOfAKind(hand string) bool {

	handArray := strings.Split(hand, "")
	sort.Strings(handArray)
	for i := 0; i < len(handArray)-3; i++ {
		if handArray[i] == handArray[i+1] && handArray[i+1] == handArray[i+2] && handArray[i+2] == handArray[i+3] {
			return true
		}
	}
	return false
}

func isFullHouse(hand string) bool {
	handArray := strings.Split(hand, "")
	sort.Strings(handArray)

	if handArray[0] == handArray[1] && handArray[1] == handArray[2] && handArray[3] == handArray[4] {
		return true
	}

	if handArray[0] == handArray[1] && handArray[2] == handArray[3] && handArray[3] == handArray[4] {
		return true
	}

	return false
}

func isThreeOfAKind(hand string) bool {

	handArray := strings.Split(hand, "")
	sort.Strings(handArray)
	for i := 0; i < len(handArray)-2; i++ {
		if handArray[i] == handArray[i+1] && handArray[i+1] == handArray[i+2] {
			return true
		}
	}
	return false
}

func isTwoPair(hand string) bool {
	handArray := strings.Split(hand, "")
	sort.Strings(handArray)

	if handArray[0] == handArray[1] && handArray[2] == handArray[3] {
		return true
	}

	if handArray[0] == handArray[1] && handArray[3] == handArray[4] {
		return true
	}

	if handArray[1] == handArray[2] && handArray[3] == handArray[4] {
		return true
	}

	return false
}

func isPair(hand string) bool {
	handArray := strings.Split(hand, "")
	sort.Strings(handArray)

	for i := 0; i < len(handArray)-1; i++ {
		if handArray[i] == handArray[i+1] {
			return true
		}
	}
	return false
}

func isHighCard(hand string) bool {
	handArray := strings.Split(hand, "")
	sort.Strings(handArray)

	if handArray[0] != handArray[1] && handArray[1] != handArray[2] && handArray[2] != handArray[3] && handArray[3] != handArray[4] {
		return true
	}
	return false
}

func sortGames(games []game) {
	sort.Slice(games, func(i, j int) bool {
		if games[i].rank == games[j].rank {
			if games[i].letterRanks[0] == games[j].letterRanks[0] {
				if games[i].letterRanks[1] == games[j].letterRanks[1] {
					if games[i].letterRanks[2] == games[j].letterRanks[2] {
						if games[i].letterRanks[3] == games[j].letterRanks[3] {
							return games[i].letterRanks[4] < games[j].letterRanks[4]
						}
						return games[i].letterRanks[3] < games[j].letterRanks[3]
					}
					return games[i].letterRanks[2] < games[j].letterRanks[2]
				}
				return games[i].letterRanks[1] < games[j].letterRanks[1]
			}
			return games[i].letterRanks[0] < games[j].letterRanks[0]
		}
		return games[i].rank < games[j].rank
	})
}

func getWinnings(games []game) int {
	totalWinnings := 0
	counter := 1
	for _, game := range games {
		totalWinnings += game.bid * counter
		counter++
	}
	return totalWinnings
}
