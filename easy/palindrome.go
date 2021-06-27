package easy

import "math"

func PalindromeNumber(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	if x < math.MinInt32 || x > math.MaxInt32 {
		return false
	}

	var reverse int

	for x > reverse {
		reverse = (reverse * 10) + (x % 10)

		x = x / 10
	}

	if reverse == x || reverse/10 == x {
		return true
	}

	return false
}
