package day12

const dot = '.'
const hash = '#'
const questionMark = '?'

func count(sequence string, numbers []int) int {
	var dp [][]int
	for i, _ := range sequence {
		dp = append(dp, make([]int, len(numbers)+1))
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}

	return calculate(0, 0, sequence, numbers, dp)
}

func calculate(i, j int, sequence string, numbers []int, options [][]int) int {
	if i >= len(sequence) {
		if j < len(numbers) {
			return 0
		}
		return 1
	}

	if options[i][j] != -1 {
		return options[i][j]
	}

	result := 0
	if sequence[i] == dot {
		result = calculate(i+1, j, sequence, numbers, options)
		options[i][j] = result
		return result
	}

	if sequence[i] == questionMark {
		result += calculate(i+1, j, sequence, numbers, options)
	}

	if j < len(numbers) {
		count := 0
		for k := i; k < len(sequence); k++ {
			if count > numbers[j] {
				break

			}
			if sequence[k] == dot {
				break
			}

			if count == numbers[j] && sequence[k] == questionMark {
				break
			}

			count++
		}

		if count == numbers[j] {
			if i+count < len(sequence) && sequence[i+count] != hash {
				result += calculate(i+count+1, j+1, sequence, numbers, options)
			} else {
				result += calculate(i+count, j+1, sequence, numbers, options)
			}
		}
	}

	options[i][j] = result
	return result
}
