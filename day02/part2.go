package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day02/input.txt")

	sumOfPower := 0

	for i := 0; i < len(lines); i++ {
		minRedCubes := 0
		minGreenCubes := 0
		minBlueCubes := 0

		line := lines[i]
		game := strings.Split(line, ":")[1]
		rounds := strings.Split(game, ";")

		for j := 0; j < len(rounds); j++ {
			round := rounds[j]
			players := strings.Split(round, ",")

			for k := 0; k < len(players); k++ {

				players[k] = strings.TrimLeft(players[k], " ")
				player := strings.Split(players[k], " ")

				number, numberErr := strconv.Atoi(player[0])
				color := strings.Trim(player[1], " ")

				if numberErr != nil {
					fmt.Println("Error:", numberErr)
					return
				}

				if color == "red" && number > minRedCubes {
					minRedCubes = number
				} else if color == "green" && number > minGreenCubes {
					minGreenCubes = number
				} else if color == "blue" && number > minBlueCubes {
					minBlueCubes = number
				}
			}
		}

		currentPower := minRedCubes * minGreenCubes * minBlueCubes
		sumOfPower += currentPower
	}

	fmt.Println("Day 2, Part 2:", sumOfPower)

}
