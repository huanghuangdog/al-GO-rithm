package other

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}

	if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	left := 0
	var tail *ListNode
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		total := v1 + v2 + left
		if head == nil {
			head = &ListNode{Val: total % 10}
			tail = head
			left = total / 10
		} else {
			tail.Next = &ListNode{Val: total % 10}
			tail = tail.Next
			left = total / 10
		}
	}
	if left > 0 {
		tail.Next = &ListNode{Val: left}
	}
	return
}
