package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	num2Map := map[int]int{}
	s := []int{}

	// example [1,3,4,2] => [3 4 -1 -1]  => {1:3, 3:4, 4:-1, 2, -1}
	for i := n - 1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= nums2[i] {
			s = s[:len(s)-1]
		}

		if len(s) == 0 {
			num2Map[nums2[i]] = -1
		} else {
			num2Map[nums2[i]] = s[len(s)-1]
		}
		s = append(s, nums2[i])
	}

	// [4,1,2], {1:3, 3:4, 4:-1, 2, -1} => [-1, 3, -1]
	answer := []int{}
	for _, val := range nums1 {
		answer = append(answer, num2Map[val])
	}
	return answer
}
