package linkedlist

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := head
	cur := head.Next
	var next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	head.Next = nil
	return pre
}
