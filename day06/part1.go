package day06

import (
	"fmt"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day06/input.txt")

	raceTimesStrings := strings.Fields(lines[0])
	winningDistancesStrings := strings.Fields(lines[1])

	raceTimes := utils.SliceAtoi(raceTimesStrings)
	winningDistances := utils.SliceAtoi(winningDistancesStrings)

	waysOfWinning := []int{}

	for i, raceTime := range raceTimes {

		prediction := 0
		currentDistance := 0

		remainingRaceTime := raceTime
		middleOfRace := raceTime - raceTime/2

		for remainingRaceTime > middleOfRace {
			remainingRaceTime--
			currentDistance++

			prediction = remainingRaceTime * currentDistance

			winningDistance := winningDistances[i]

			if prediction > winningDistance {
				wayOfWinning := raceTime - currentDistance*2 + 1
				waysOfWinning = append(waysOfWinning, wayOfWinning)
				break
			}
		}
	}

	total := utils.MultiplyArray(waysOfWinning)

	fmt.Println("Day 6, Part 1:", total)

}
