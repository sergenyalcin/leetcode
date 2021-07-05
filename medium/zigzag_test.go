package medium

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func TestConvert(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		str     string
		numRows int
		output  string
	}{
		{"AB", 1, "AB"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := Convert(tc.str, tc.numRows)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
