package main

func maximalSquare(matrix [][]byte) int {
	m := len(matrix[0])
	n := len(matrix)
	lengths := make([][]int, n)
	for i := range lengths {
		lengths[i] = make([]int, m)
	}
	maxArea := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				up, diag, left := 0, 0, 0
				if i > 0 {
					up = lengths[i-1][j]
				}

				if j > 0 {
					left = lengths[i][j-1]
				}

				if i > 0 && j > 0 {
					diag = lengths[i-1][j-1]
				}

				lengths[i][j] = min(min(up, diag), left) + 1
				maxArea = max(maxArea, lengths[i][j]*lengths[i][j])
			}
		}
	}

	return maxArea
}
