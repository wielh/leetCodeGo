package main

func moveZeroes(nums []int) {
	shift := 0
	for index, num := range nums {
		if num == 0 {
			shift += 1
		} else if shift > 0 {
			tmp := num
			nums[index] = nums[index-shift]
			nums[index-shift] = tmp
		}
	}
}
