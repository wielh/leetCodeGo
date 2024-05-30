package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	charMap := map[byte]int{}
	currentMaxLen := 0

	left := 0
	right := 0

	for right < len(s) {
		char := s[right]
		index, ok := charMap[char]
		if ok && index+1 > left {
			left = index + 1
		}

		charMap[char] = right
		if right-left+1 > currentMaxLen {
			currentMaxLen = right - left + 1
		}
		right++
	}
	return currentMaxLen
}
