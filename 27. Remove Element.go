package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}

	leftIndex := 0
	rightIndex := len(nums) - 1
	answer := 0

	for {
		if nums[leftIndex] != val {
			leftIndex += 1
			answer += 1
			if leftIndex > rightIndex {
				return answer
			}
		} else {
			for nums[rightIndex] == val {
				rightIndex -= 1
				if leftIndex > rightIndex {
					return answer
				}
			}
			tmp := nums[rightIndex]
			nums[rightIndex] = nums[leftIndex]
			nums[leftIndex] = tmp
		}
	}
}
