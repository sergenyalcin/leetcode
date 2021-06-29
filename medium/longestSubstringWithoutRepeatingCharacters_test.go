package medium

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		str    string
		output int
	}{
		{"abba", 2},
		{"dvdf", 3},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := LengthOfLongestSubstring(tc.str)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
