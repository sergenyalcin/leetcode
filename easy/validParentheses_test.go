package easy

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	//"github.com/onsi/gomega"
)

func TestValidParentheses(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		str    string
		output bool
	}{
		{"(", false},
		{")", false},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := ValidParentheses(tc.str)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
