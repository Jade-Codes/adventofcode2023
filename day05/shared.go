package day05

import (
	"strconv"
	"strings"
)

type category struct {
	destinationRangeStart, destinationRangeEnd, sourceRangeStart, sourceRangeEnd, rangeCount int
}

func addToCategories(line string, currentCategories []category) []category {
	currentNumbers := strings.Fields(line)

	destinationRangeStart, _ := strconv.Atoi(currentNumbers[0])
	sourceRangeStart, _ := strconv.Atoi(currentNumbers[1])
	rangeCount, _ := strconv.Atoi(currentNumbers[2])

	currentCategory := category{
		destinationRangeStart: destinationRangeStart,
		destinationRangeEnd:   destinationRangeStart + rangeCount - 1,
		sourceRangeStart:      sourceRangeStart,
		sourceRangeEnd:        sourceRangeStart + rangeCount - 1,
		rangeCount:            rangeCount,
	}

	return append(currentCategories, currentCategory)
}
