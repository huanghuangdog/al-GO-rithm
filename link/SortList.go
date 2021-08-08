package link

func sortList(head *ListNode) *ListNode {
	return sort(head)
}

func sort(head *ListNode) *ListNode {
	one, two := divide(head)
	if one == two {
		return one
	} else {
		one = sort(one)
		two = sort(two)
		return merge(one, two)
	}
}

func merge(one *ListNode, two *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp, tmp1, tmp2 := dummyHead, one, two

	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val <= tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	for tmp1 != nil {
		tmp.Next = tmp1
		tmp1 = tmp1.Next
	}
	for tmp2 != nil {
		tmp.Next = tmp2
		tmp2 = tmp2.Next
	}
	return dummyHead.Next
}

func divide(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	prevSlow, slow, fast := head, head, head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil {
		tail := slow.Next
		slow.Next = nil
		return head, tail
	} else {
		prevSlow.Next = nil
		return head, slow
	}
}
