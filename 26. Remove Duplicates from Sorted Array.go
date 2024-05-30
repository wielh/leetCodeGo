package main

func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	arrayLength := len(nums)
	leftIndex := 0
	rightIndex := 0
	shift := 0
	answer := 0

	for {
		//fmt.Println(leftIndex, rightIndex, shift, answer, nums)
		if rightIndex >= arrayLength-shift {
			return answer
		}

		for rightIndex+1 < arrayLength-shift && nums[leftIndex] == nums[rightIndex+1] {
			rightIndex += 1
		}

		swap(nums, leftIndex, rightIndex)
		shift += (rightIndex - leftIndex)
		leftIndex += 1
		rightIndex = leftIndex
		answer += 1
	}
}

func swap(nums []int, leftIndex int, rightIndex int) {
	// [leftIndex+1, rightIndex]
	if leftIndex == rightIndex {
		return
	}

	length := rightIndex - leftIndex
	for i := rightIndex + 1; i < len(nums); i++ {
		nums[i-length] = nums[i]
	}
}
