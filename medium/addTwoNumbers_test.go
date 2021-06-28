package medium

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

var node1 = ListNode{Val: 9, Next: &node2}
var node2 = ListNode{Val: 9, Next: &node3}
var node3 = ListNode{Val: 9, Next: &node4}
var node4 = ListNode{Val: 9, Next: &node5}
var node5 = ListNode{Val: 9, Next: &node6}
var node6 = ListNode{Val: 9, Next: &node7}
var node7 = ListNode{Val: 9, Next: nil}

var node8 = ListNode{Val: 9, Next: &node9}
var node9 = ListNode{Val: 9, Next: &node10}
var node10 = ListNode{Val: 9, Next: &node11}
var node11 = ListNode{Val: 9, Next: nil}

var out1 = ListNode{Val: 8, Next: &out2}
var out2 = ListNode{Val: 9, Next: &out3}
var out3 = ListNode{Val: 9, Next: &out4}
var out4 = ListNode{Val: 9, Next: &out5}
var out5 = ListNode{Val: 0, Next: &out6}
var out6 = ListNode{Val: 0, Next: &out7}
var out7 = ListNode{Val: 0, Next: &out8}
var out8 = ListNode{Val: 1, Next: nil}

func TestAddTwoNumbers(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		l1     *ListNode
		l2     *ListNode
		output *ListNode
	}{
		{&node1, &node8, &out1},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := AddTwoNumbers(&node1, &node4)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
