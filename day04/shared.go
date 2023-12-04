package day04

func getNumbersWon(winningNumbersArray []string, playingNumbersArray []string) int {
	numbersWon := 0

	for _, number := range winningNumbersArray {
		for _, playingNumber := range playingNumbersArray {
			if number == playingNumber {
				numbersWon++
			}
		}
	}

	return numbersWon
}
