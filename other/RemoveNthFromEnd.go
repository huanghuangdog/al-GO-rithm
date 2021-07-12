package other

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	step := 0
	for fast.Next != nil && step < n {
		fast = fast.Next
		step++
	}
	if step < n {
		return slow.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
