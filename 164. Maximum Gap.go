package main

import "sort"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)
	answer := 0
	currentNum := 0
	for i := 0; i < len(nums)-1; i++ {
		currentNum = nums[i+1] - nums[i]
		if currentNum > answer {
			answer = currentNum
		}
	}
	return answer
}
