package easy

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	//"github.com/onsi/gomega"
)

var node1 = ListNode{Val: 1, Next: &node2}
var node2 = ListNode{Val: 2, Next: &node3}
var node3 = ListNode{Val: 4, Next: nil}

var node4 = ListNode{Val: 1, Next: &node5}
var node5 = ListNode{Val: 3, Next: &node6}
var node6 = ListNode{Val: 4, Next: nil}

var out1 = ListNode{Val: 1, Next: &out2}
var out2 = ListNode{Val: 1, Next: &out3}
var out3 = ListNode{Val: 2, Next: &out4}
var out4 = ListNode{Val: 3, Next: &out5}
var out5 = ListNode{Val: 4, Next: &out6}
var out6 = ListNode{Val: 4, Next: nil}

func TestMergeTwoSortedLists(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testCases := []struct {
		l1     *ListNode
		l2     *ListNode
		output *ListNode
	}{
		{&node1, &node4, &out1},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
			acutal := MergeTwoSortedLists(tc.l1, tc.l2)

			g.Expect(acutal).To(gomega.Equal(tc.output))
		})
	}
}
