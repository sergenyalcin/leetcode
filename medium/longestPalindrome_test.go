package medium

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func TestLongestPalindrome(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		str    string
		output string
	}{
		{"aacabdkacaa", "aca"},
		{"aaaa", "aaaa"},
		{"cbbd", "bb"},
		{"babad", "bab"},
		{"a", "a"},
		{"ac", "a"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := LongestPalindrome(tc.str)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
