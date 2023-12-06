package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetLines(filename string) []string {
	filePath := filename

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var s []string

	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	return s
}

func SliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err == nil {
			si = append(si, i)
		}
	}
	return si
}

func MultiplyArray(array []int) int {
	total := 1
	for _, number := range array {
		total *= number
	}
	return total
}
