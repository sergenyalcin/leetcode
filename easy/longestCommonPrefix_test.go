package easy

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	//"github.com/onsi/gomega"
)

func TestLongestCommonPrefix(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		strs   []string
		output string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := LongestCommonPrefix(tc.strs)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
