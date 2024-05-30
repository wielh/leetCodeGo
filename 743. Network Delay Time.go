package main

func networkDelayTime(times [][]int, n int, k int) int {

	type Node struct {
		Vertex int
		Time   int
	}

	// sort times by start point
	timesMap := map[int][]Node{}
	for _, time := range times {
		currentMap, ok := timesMap[time[0]]
		if ok {
			timesMap[time[0]] = append(currentMap, Node{Vertex: time[1], Time: time[2]})
		} else {
			timesMap[time[0]] = []Node{{Vertex: time[1], Time: time[2]}}
		}
	}

	// start calculate
	currentPoints := []int{k}
	nextPoints := []int{}
	currentArriveTimeMap := map[int]int{k: 0}

	for len(currentPoints) > 0 {
		for _, src := range currentPoints {
			paths, ok := timesMap[src]
			if !ok {
				continue
			}

			srcArriveTime := currentArriveTimeMap[src]
			for _, path := range paths {
				destArriveTime, destOk := currentArriveTimeMap[path.Vertex]
				if !destOk || srcArriveTime+path.Time < destArriveTime {
					currentArriveTimeMap[path.Vertex] = path.Time + srcArriveTime
					nextPoints = append(nextPoints, path.Vertex)
				}
			}
		}

		currentPoints = nextPoints
		nextPoints = []int{}
	}

	if len(currentArriveTimeMap) < n {
		return -1
	}

	max := 0
	for _, value := range currentArriveTimeMap {
		if max < value {
			max = value
		}
	}
	return max
}
