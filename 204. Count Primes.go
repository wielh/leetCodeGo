package main

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if !isNotPrime[i] {
			for j := i * i; j < n; j += i {
				isNotPrime[j] = true
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			count++
		}
	}

	return count
}
