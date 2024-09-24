package main

func longestNiceSubarray(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	maxLen := 0
	currentLen := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]&nums[i+1] == 0 {
			currentLen += 1
		} else {
			maxLen = max(maxLen, currentLen)
			currentLen = 0
		}
	}
	return maxLen
}
