package day17

import (
	"container/heap"
	"fmt"
	"strconv"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day17/input.txt")
	matrix := [][]int{}
	for y, line := range lines {
		matrix = append(matrix, []int{})
		for _, char := range line {
			matrixInt, _ := strconv.Atoi(string(char))
			matrix[y] = append(matrix[y], matrixInt)
		}
	}

	heatLoss := findLowestHeatLossWithUltra(matrix)

	fmt.Println("Day 17, Part 1:", heatLoss)
}

func findLowestHeatLossWithUltra(matrix [][]int) int {
	maxX := len(matrix[0]) - 1
	maxY := len(matrix) - 1

	locationValuesMap := make(map[Location]int)

	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)

	start := Location{0, 0, east, 0}

	heap.Push(&priorityQueue, struct {
		location Location
		dist     int
	}{start, 0})

	for priorityQueue.Len() > 0 {
		current := heap.Pop(&priorityQueue).(struct {
			location Location
			dist     int
		})

		x, y := current.location.x, current.location.y
		currentHeatLoss := current.dist

		if locationValuesMap[current.location] < currentHeatLoss && locationValuesMap[current.location] != 0 {
			continue
		}

		locationValuesMap[current.location] = currentHeatLoss

		if current.location.x == maxX && current.location.y == maxY {
			return currentHeatLoss
		}

		locations := getNextLocationsUltra(x, y, current.location.direction, current.location.directionCount)

		for _, location := range locations {
			if location.x < 0 || location.y < 0 || location.x > maxX || location.y > maxY {
				continue
			}
			newHeatLoss := currentHeatLoss + matrix[location.y][location.x]

			if dist, found := locationValuesMap[location]; !found || newHeatLoss < dist {
				locationValuesMap[location] = newHeatLoss
				heap.Push(&priorityQueue, struct {
					location Location
					dist     int
				}{location, newHeatLoss})
			}
		}
	}

	return -1
}

func getNextLocationsUltra(x int, y int, direction rune, currentCount int) []Location {
	locations := []Location{}
	if currentCount < 10 {
		locations = append(locations, getNextLocationUltra(x, y, direction, currentCount+1))
	}
	if currentCount >= 4 {
		switch direction {
		case north:
			locations = append(locations, getNextLocationUltra(x, y, east, 1))
			locations = append(locations, getNextLocationUltra(x, y, west, 1))
		case south:
			locations = append(locations, getNextLocationUltra(x, y, east, 1))
			locations = append(locations, getNextLocationUltra(x, y, west, 1))
		case east:
			locations = append(locations, getNextLocationUltra(x, y, north, 1))
			locations = append(locations, getNextLocationUltra(x, y, south, 1))
		case west:
			locations = append(locations, getNextLocationUltra(x, y, north, 1))
			locations = append(locations, getNextLocationUltra(x, y, south, 1))
		}
	}
	return locations
}

func getNextLocationUltra(x int, y int, direction rune, count int) Location {
	switch direction {
	case north:
		y--
	case south:
		y++
	case east:
		x++
	case west:
		x--
	}
	return Location{x, y, direction, count}
}
