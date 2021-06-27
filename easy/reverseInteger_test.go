package easy

import (
	"fmt"
	"math"
	"testing"

	"github.com/onsi/gomega"
	//"github.com/onsi/gomega"
)

func TestReverseInteger(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		num    int
		output int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{math.MaxInt32 + 100, 0},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := ReverseInteger(tc.num)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
