package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part1() {
	lines := utils.GetLines("day02/input.txt")

	redCubes := 12
	greenCubes := 13
	blueCubes := 14

	successfulGames := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		game := strings.Split(line, ":")[1]
		rounds := strings.Split(game, ";")

		isSuccessful := true
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

				if color == "red" && number > redCubes {
					isSuccessful = false
					break
				} else if color == "green" && number > greenCubes {
					isSuccessful = false
					break
				} else if color == "blue" && number > blueCubes {
					isSuccessful = false
					break
				}
			}
			if !isSuccessful {
				break
			}

		}
		if isSuccessful {
			successfulGames += i + 1
		}

	}

	fmt.Println("Day 2, Part 1:", successfulGames)

}
