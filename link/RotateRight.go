package link

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 || head.Next == nil{
		return nil
	}
	slow, fast := head, head
	for i:= 0; i < k;i++ {
		if fast.Next == nil {
			fast = head
		} else {
			fast = fast.Next
		}
	}
	if slow == fast {
		return head
	}
	for slow.Next != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	newHead := slow.Next
	fast.Next = head
	slow.Next = nil
	return newHead
}