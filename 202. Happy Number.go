package main

func isHappy(n int) bool {
	return isHappyNumber(n, n)
}

func isHappyNumber(n int, initNum int) bool {
	if n == 1 {
		return true
	}
	m := sumOfSquare(n)
	if m == initNum {
		return false
	}
	return isHappyNumber(m, initNum)
}

func sumOfSquare(n int) int {
	sum := 0
	d := 0
	for {
		d = n % 10
		n = (n - d) / 10
		sum += d * d
		if n == 0 {
			break
		}
	}
	return sum
}
