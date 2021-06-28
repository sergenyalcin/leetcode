package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int

	result := new(ListNode)
	current := result

	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = &ListNode{
				Val:  0,
				Next: nil,
			}
		} else if l2 == nil {
			l2 = &ListNode{
				Val:  0,
				Next: nil,
			}
		}

		sum := l1.Val + l2.Val + carry

		carry = sum / 10

		current.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if carry > 0 {
		current.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return result.Next
}
