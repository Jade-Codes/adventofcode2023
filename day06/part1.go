package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day06/input.txt")

	raceTimesStrings := strings.Fields(lines[0])
	winningDistancesStrings := strings.Fields(lines[1])

	raceTimes := sliceAtoi(raceTimesStrings)
	winningDistances := sliceAtoi(winningDistancesStrings)

	waysOfWinning := []int{}

	for i, raceTime := range raceTimes {
		previousTime := 0
		currentDistance := 0
		remainingRaceTime := raceTime

		middleOfRace := raceTime - raceTime/2

		for remainingRaceTime > middleOfRace {

			remainingRaceTime--

			previousTime = remainingRaceTime * (currentDistance + 1)

			currentDistance++
			winningDistance := winningDistances[i]

			if previousTime > winningDistance {
				wayOfWinning := raceTime - currentDistance*2 + 1
				waysOfWinning = append(waysOfWinning, wayOfWinning)
				break
			}
		}
	}

	total := 1

	for _, wayOfWinning := range waysOfWinning {
		total *= wayOfWinning
	}

	fmt.Println("Day 6, Part 1:", total)

}

func sliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err == nil {
			si = append(si, i)
		}
	}
	return si
}
