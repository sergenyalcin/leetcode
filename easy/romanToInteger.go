package easy

func RomanToInteger(s string) int {
	romanNumbers := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var currentValue, prevValue, total int

	for _, c := range s {
		currentValue = romanNumbers[c]

		if currentValue > prevValue && prevValue > 0 {
			currentValue -= 2 * prevValue
		}

		total += currentValue

		prevValue = currentValue
	}

	return total
}
