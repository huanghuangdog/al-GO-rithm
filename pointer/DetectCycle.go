package pointer

type ListNode struct {
	Val int
	Next *ListNode
}
/**
20210815
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil ||fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
