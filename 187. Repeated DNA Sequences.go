package main

func findRepeatedDnaSequences(s string) []string {
	answer := []string{}
	if len(s) <= 10 {
		return answer
	}

	appearTimes := map[string]int{}
	for i := 0; i < len(s)-10; i++ {
		appearTimes[s[i:i+10]]++
	}

	for key, times := range appearTimes {
		if times > 1 {
			answer = append(answer, key)
		}
	}
	return answer
}
