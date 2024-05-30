package main

import "fmt"

type stringStatus struct {
	initRune   []rune
	leftIndex  int
	rightIndex int
	charCount  map[rune]int
	currentMax int
	currentMin int
}

func (s *stringStatus) init(initStr string) {
	s.initRune = []rune(initStr)
	s.leftIndex = -1
	s.leftBackward()
}

func (s *stringStatus) leftBackward() {
	if s.leftIndex >= len(s.initRune) {
		return
	}
	s.leftIndex += 1
	s.rightIndex = s.leftIndex
	s.charCount = map[rune]int{}
	s.currentMin = 1000
	s.currentMax = 0
}

func (s *stringStatus) rightForward() {
	if s.rightIndex >= len(s.initRune) {
		return
	}

	rightChar := s.initRune[s.rightIndex]
	val := s.charCount[rightChar]
	val++
	s.charCount[rightChar] = val
	if val > s.currentMax {
		s.currentMax = val
	}

	if s.currentMin+1 >= val {
		s.currentMin = 1000
		for _, val := range s.charCount {
			if val < s.currentMin {
				s.currentMin = val
			}
		}
	}
	s.rightIndex++
}

func (s *stringStatus) calculateBeauty() int {
	fmt.Println(s.charCount, s.currentMax-s.currentMin)
	return s.currentMax - s.currentMin
}

func beautySum(s string) int {
	sum := 0
	status := stringStatus{}
	status.init(s)
	len := len(status.initRune)
	for status.leftIndex < len {
		for status.rightIndex < len {
			status.rightForward()
			sum += status.calculateBeauty()
		}
		status.leftBackward()
	}
	return sum
}
