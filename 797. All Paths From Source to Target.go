package main

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	currentTmpAnswer := [][]int{{0}}
	currentAnswer := [][]int{}
	nextTmpAnswer := [][]int{}

	for {
		for _, path := range currentTmpAnswer {
			lastPoint := path[len(path)-1]
			nextPath := make([]int, len(path))
			copy(nextPath, path)

			for _, nextPoint := range graph[lastPoint] {
				if nextPoint == n-1 {
					currentAnswer = append(currentAnswer, append(nextPath, nextPoint))
				} else {
					nextTmpAnswer = append(nextTmpAnswer, append(nextPath, nextPoint))
				}
			}
		}

		if len(nextTmpAnswer) == 0 {
			return currentAnswer
		}
		currentTmpAnswer = nextTmpAnswer
		nextTmpAnswer = [][]int{}
	}
}
