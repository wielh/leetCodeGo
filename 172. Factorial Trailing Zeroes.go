package main

func trailingZeroes(n int) int {
	digits := []int{}
	for {
		mod := n % 5
		digits = append(digits, mod)
		n = (n - mod) / 5
		if n <= 0 {
			break
		}
	}

	answer := 0
	zeroes := 0
	for _, digit := range digits {
		answer += zeroes * digit
		zeroes = 5*zeroes + 1
	}

	return answer
}
