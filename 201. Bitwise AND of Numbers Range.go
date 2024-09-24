package main

func rangeBitwiseAnd(left int, right int) int {
	numBits := 32
	mix := left & right
	for i := numBits - 1; i >= 0; i-- {
		bit := (mix >> i) & 1
		if bit == 0 {
			if i == 31 {
				return 0
			} else {
				return bit << (i + 1)
			}
		}
	}
	return left
}
