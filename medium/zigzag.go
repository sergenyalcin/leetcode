package medium

func Convert(s string, numRows int) string {
	word := make([][]byte, numRows)
	result := ""
	i := 0
	currentRow := -1

	if numRows == 1 {
		return s
	}

	inc := func(j *int, amount int) {
		*j += amount
	}

	dec := func(j *int, amount int) {
		*j -= amount
	}

	f := inc

	for i < len(s) {
		f(&currentRow, 1)
		word[currentRow] = append(word[currentRow], s[i])

		if currentRow == 0 {
			f = inc
		}

		if currentRow == numRows-1 {
			f = dec
		}

		i++
	}

	for _, row := range word {
		result += string(row)
	}

	return result
}
