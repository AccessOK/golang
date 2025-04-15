package main

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil && pHead2 == nil {
		return nil
	}
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	p1 := pHead1
	p2 := pHead2
	var result *ListNode
	if p1.Val < p2.Val {
		result = pHead1
		p1 = pHead1.Next
	} else {
		result = pHead2
		p2 = pHead2.Next
	}
	p := result
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p = p.Next
			p1 = p1.Next
		} else {
			p.Next = p2
			p = p.Next
			p2 = p2.Next
		}
	}
	for p1 != nil {
		p.Next = p1
		p1 = p1.Next
		p = p.Next
	}
	for p2 != nil {
		p.Next = p2
		p2 = p2.Next
		p = p.Next
	}
	return result
}
