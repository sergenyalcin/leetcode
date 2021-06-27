package easy

import "math"

func ReverseInteger(x int) int {
	var reverse int

	for x != 0 {
		reverse = (reverse * 10) + (x % 10)

		x = x / 10
	}

	if reverse < math.MinInt32 || reverse > math.MaxInt32 {
		return 0
	}

	return reverse
}
