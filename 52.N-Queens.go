package main

func totalNQueens(n int) int {
	return placeQueens2(n, [][]int{})
}

func placeQueens2(n int, currentQueensLocation [][]int) int {
	if len(currentQueensLocation) == n {
		return 1
	}

	row := len(currentQueensLocation)
	answer := 0
	for col := 0; col < n; col++ {
		ok := true
		for _, queensLocation := range currentQueensLocation {
			if col == queensLocation[1] ||
				row+col == queensLocation[0]+queensLocation[1] ||
				row-col == queensLocation[0]-queensLocation[1] {
				ok = false
				break
			}
		}

		if ok {
			nextQueensLocation := append([][]int{{row, col}}, currentQueensLocation...)
			answer += placeQueens2(n, nextQueensLocation)
		}
	}
	return answer
}
