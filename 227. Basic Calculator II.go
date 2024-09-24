package main

func calculate(s string) int {
	currentNum := 0
	nums := []int{}
	operators := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			currentNum = 10*currentNum + int(s[i]-'0')
		case '+', '-', '*', '/':
			nums = append(nums, currentNum)
			currentNum = 0
			operators = append(operators, s[i])
		}
	}
	nums = append(nums, currentNum)

	nums2 := []int{}
	operators2 := []byte{}
	currentNum = nums[0]
	for i, op := range operators {
		switch op {
		case '+', '-':
			nums2 = append(nums2, currentNum)
			currentNum = nums[i+1]
			operators2 = append(operators2, op)
		case '*':
			currentNum = currentNum * nums[i+1]
		case '/':
			currentNum = currentNum / nums[i+1]
		}
	}
	nums2 = append(nums2, currentNum)

	answer := nums2[0]
	for i, op := range operators2 {
		switch op {
		case '+':
			answer += nums2[i+1]
		case '-':
			answer -= nums2[i+1]
		}
	}
	return answer
}
