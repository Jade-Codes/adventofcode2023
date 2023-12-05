package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day06/input.txt")

	raceTimeString := strings.Replace(lines[0], "Time:", "", -1)
	raceTimeString = strings.Replace(raceTimeString, " ", "", -1)

	winningDistanceString := strings.Replace(lines[1], "Distance:", "", -1)
	winningDistanceString = strings.Replace(winningDistanceString, " ", "", -1)

	raceTime, _ := strconv.Atoi(raceTimeString)
	winningDistance, _ := strconv.Atoi(winningDistanceString)

	wayOfWinning := 0

	previousTime := 0
	currentDistance := 0
	remainingRaceTime := raceTime

	middleOfRace := raceTime - raceTime/2

	for remainingRaceTime > middleOfRace {

		remainingRaceTime--

		previousTime = remainingRaceTime * (currentDistance + 1)

		currentDistance++
		winningDistance := winningDistance

		if previousTime > winningDistance {
			wayOfWinning = raceTime - currentDistance*2 + 1
			break
		}
	}

	fmt.Println("Day 6, Part 1:", wayOfWinning)

}

func multiplyArray(ia []int) int {

	total := 1

	for _, i := range ia {
		total *= i
	}

	return total
}
