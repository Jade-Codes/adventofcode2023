package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part2() {
	filePath := "day1/day1.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() //

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile("((\\d)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine))")
		re2 := regexp.MustCompile("((\\d)|(eno)|(owt)|(eerht)|(ruof)|(evif)|(xis)|(neves)|(thgie)|(enin))")

		match1 := re.FindStringSubmatch(line)
		match2 := re2.FindStringSubmatch(reverse(line))

		if match1 != nil && match2 != nil {
			number1 := match1[0]
			number2 := reverse(match2[0])

			if len(number1) > 1 {

				number1 = digitMap[number1]
			}
			if len(number2) > 1 {
				number2 = digitMap[number2]
			}
			numberString := number1 + number2
			number, numberErr := strconv.Atoi(numberString)

			if numberErr != nil {
				fmt.Println("Error:", numberErr)
				return
			}

			sum += number
		}
	}

	fmt.Println("Day 1, Part 2:", sum)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
