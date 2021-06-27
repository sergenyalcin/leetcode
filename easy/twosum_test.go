package easy

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func TestSum(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		nums   []int
		target int
		output []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := TwoSum(tc.nums, tc.target)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
