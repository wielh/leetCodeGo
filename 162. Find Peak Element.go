package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums)
}
