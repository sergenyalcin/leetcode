package easy

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	//"github.com/onsi/gomega"
)

func TestRomanToInteger(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		roman  string
		output int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := RomanToInteger(tc.roman)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
