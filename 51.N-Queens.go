package main

import "strings"

func solveNQueens(n int) [][]string {
	answer := [][]string{}
	placeQueens(n, &[][]string{}, &answer)
	return answer
}

func placeQueens(n int, board *[][]string, answer *[][]string) {
	if len(*board) == n {
		*answer = append(*answer, locationsToStringArray(*board))
		return
	}

	row := len(*board)
	for col := 0; col < n; col++ {
		if judge(row, col, *board) {
			*board = append(*board, colToStringArray(n, col))
			placeQueens(n, board, answer)
			b := *board
			*board = b[:len(*board)-1]
		}
	}
}

func colToStringArray(n int, col int) []string {
	answer := []string{}
	for i := 0; i < n; i++ {
		if i == col {
			answer = append(answer, "Q")
			continue
		}
		answer = append(answer, ".")
	}
	return answer
}

func judge(row int, col int, currentQueensLocation [][]string) bool {
	for r, rowStrs := range currentQueensLocation {
		for c, str := range rowStrs {
			if str == "Q" {
				if c == col || r+c == row+col || r-c == row-col {
					return false
				}
				break
			}
		}
	}
	return true
}

func locationsToStringArray(locations [][]string) []string {
	answers := []string{}
	for _, location := range locations {
		answers = append(answers, strings.Join(location, ""))
	}
	return answers
}
