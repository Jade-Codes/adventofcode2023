package day01

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Jade-Codes/adventofcode2023/utils"
)

func Part2() {
	lines := utils.GetLines("day01/input.txt")

	sum := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
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
