package medium

func LongestPalindrome(s string) string {
	maxLen := 0
	startIndex := 0

	l := len(s)

	for i := 0; i < l; i++ {
		maxLen, startIndex = extendPalindrome(s, i, i, maxLen, startIndex)
		maxLen, startIndex = extendPalindrome(s, i, i+1, maxLen, startIndex)
	}

	if maxLen < 1 {
		return string(s[0])
	}

	return s[startIndex : startIndex+maxLen]
}

func extendPalindrome(s string, start, end, maxLen, startIndex int) (int, int) {
	for start >= 0 && end < len(s) {
		if s[start] == s[end] {
			if end-start+1 > maxLen {
				startIndex = start
				maxLen = end - start + 1
			}
		} else {
			break
		}

		start--
		end++
	}

	return maxLen, startIndex
}
