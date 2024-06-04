package main

import (
	"sort"
)

type SlidingWindow struct {
	valMap       map[int][]int
	valArrayDesc []int
}

func (s *SlidingWindow) insertVal(val int) {
	index := sort.Search(len(s.valArrayDesc), func(i int) bool {
		return s.valArrayDesc[i] <= val
	})

	s.valArrayDesc = s.valArrayDesc[:index]
	s.valArrayDesc = append(s.valArrayDesc, val)
}

func (s *SlidingWindow) init(valArray []int, size int) {
	s.valMap = map[int][]int{}
	for index, value := range valArray[:size] {
		array, ok := s.valMap[value]
		if !ok {
			s.valMap[value] = []int{index}
		} else {
			s.valMap[value] = append(array, index)
		}
	}
	s.valArrayDesc = []int{}
	for _, val := range valArray[0:size] {
		s.insertVal(val)
	}
}

func (s *SlidingWindow) operate(leftVal int, leftIndex int, rightVal int, rightIndex int) {
	// remove value: leftVal  from s.valMap
	array, ok := s.valMap[leftVal]
	if ok && len(array) > 0 && array[0] == leftIndex {
		s.valMap[leftVal] = array[1:]
	}

	// remove value: leftVal =s.valArrayDesc[0] from s.valArrayDesc
	if leftVal == s.valArrayDesc[0] {
		array, ok := s.valMap[leftVal]
		if ok && len(array) == 0 {
			s.valArrayDesc = s.valArrayDesc[1:]
		}
	}

	array, ok = s.valMap[rightVal]
	if !ok {
		s.valMap[rightVal] = []int{rightIndex}
	} else {
		s.valMap[rightVal] = append(array, rightIndex)
	}
	s.insertVal(rightVal)
}

func (s *SlidingWindow) getMaxValue() int {
	return s.valArrayDesc[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	window := SlidingWindow{}
	window.init(nums, k)
	answer := []int{window.getMaxValue()}
	for i := k; i < len(nums); i++ {
		window.operate(nums[i-k], i-k, nums[i], i)
		answer = append(answer, window.getMaxValue())
	}
	return answer
}
