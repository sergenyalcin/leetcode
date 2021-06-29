package medium

func LengthOfLongestSubstring(s string) int {
	chars := make(map[string]int)
	var max int
	i := 0

	for j := range s {
		if v, ok := chars[string(s[j])]; ok {
			if i < v+1 {
				i = v + 1
			}
		}

		if j-i+1 > max {
			max = j - i + 1
		}

		chars[string(s[j])] = j

		j++
	}

	return max
}
