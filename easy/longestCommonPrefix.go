package easy

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	longest := strs[0]

	for _, s := range strs[1:] {
		var common string

		for _, c := range s {
			char := string(c)

			if strings.HasPrefix(longest, common+char) {
				common += char
			} else {
				break
			}
		}

		longest = common
	}

	return longest
}
