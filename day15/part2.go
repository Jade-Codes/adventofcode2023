package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

type Box struct {
	label       string
	focalLength int
}

func Part2() {
	lines := utils.GetLines("day15/input.txt")

	boxes := make(map[int][]Box)

	for _, line := range lines {
		codes := strings.Split(line, ",")
		for _, code := range codes {
			label := ""
			focalLength := -1
			if strings.Contains(code, "-") {
				label = strings.Split(code, "-")[0]
			} else {
				label = strings.Split(code, "=")[0]
				focalLength, _ = strconv.Atoi(strings.Split(code, "=")[1])
			}

			box := calculateHash(label)

			if focalLength != -1 {
				currentBox := boxes[box]

				alreadyExists := false

				for i, b := range currentBox {
					if b.label == label {
						boxes[box][i].focalLength = focalLength
						alreadyExists = true
						break
					}
				}

				if !alreadyExists {
					boxes[box] = append(boxes[box], Box{label, focalLength})
				}
			} else {
				currentBox := boxes[box]
				for i, b := range currentBox {
					if b.label == label {
						boxes[box] = append(currentBox[:i], currentBox[i+1:]...)
						break
					}

				}
			}
		}

	}

	totalFocusingPower := calculateFocusingPower(boxes)

	fmt.Println("Day 15, Part 1:", totalFocusingPower)
}

func calculateFocusingPower(boxes map[int][]Box) int {
	totalFocusingPower := 0

	for i, box := range boxes {
		boxNumber := i + 1
		for j, box := range box {
			slotNumber := j + 1
			focalLength := box.focalLength
			currentFocusingPower := boxNumber * slotNumber * focalLength
			totalFocusingPower += currentFocusingPower
		}
	}

	return totalFocusingPower

}
