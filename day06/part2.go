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

	prediction := 0
	currentDistance := 0
	remainingRaceTime := raceTime

	middleOfRace := raceTime - raceTime/2

	for remainingRaceTime > middleOfRace {

		remainingRaceTime--
		currentDistance++

		prediction = remainingRaceTime * currentDistance
		winningDistance := winningDistance

		if prediction > winningDistance {
			wayOfWinning = raceTime - currentDistance*2 + 1
			break
		}
	}

	fmt.Println("Day 6, Part 2:", wayOfWinning)

}
