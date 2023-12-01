package utils

import (
	"bufio"
	"fmt"
	"os"
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
