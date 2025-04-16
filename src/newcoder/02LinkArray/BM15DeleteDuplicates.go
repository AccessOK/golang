package LinkArray

func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	p := head
	for p != nil && p.Next != nil {
		if p.Next.Val == p.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}
