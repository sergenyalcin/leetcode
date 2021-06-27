package easy

import (
	"fmt"
	"math"
	"testing"

	"github.com/onsi/gomega"
	//"github.com/onsi/gomega"
)

func TestPalindromeNumber(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		num    int
		output bool
	}{
		{121, true},
		{11, true},
		{123, false},
		{-121, false},
		{10, false},
		{-101, false},
		{0, true},
		{math.MaxInt32 + 100, false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := PalindromeNumber(tc.num)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
