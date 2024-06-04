package main

func maximalRectangle(matrix [][]byte) int {
	heights := make([]int, len(matrix[0])+1)
	heights[len(heights)-1] = -1
	maxArea := 0

	for _, row := range matrix {

		/* example
		10100       lst 10100
		10111       2nd 20211
		11111    => 3rd 31322
		10010       4th:40030
		*/
		for i := range row {
			if row[i] == '1' {
				heights[i]++
			} else {
				heights[i] = 0
			}
		}

		stack := []int{}
		for i, currentHeight := range heights {
			for len(stack) > 0 {
				lastElement := stack[len(stack)-1]
				if heights[lastElement] > currentHeight {
					prev := heights[lastElement]
					stack = stack[:len(stack)-1]

					width := i
					if len(stack) > 0 {
						width = i - lastElement - 1
					}
					maxArea = max(maxArea, prev*width)
				}
			}
			stack = append(stack, i)
		}
	}
	return maxArea
}
