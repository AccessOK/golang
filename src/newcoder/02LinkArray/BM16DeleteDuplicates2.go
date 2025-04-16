package LinkArray

func deleteDuplicates2(head *ListNode) *ListNode {
	// write code here
	p := head
	pre := new(ListNode)
	result := pre
	for p != nil {
		if p.Next == nil {
			pre.Next = p
			break
		}
		q := p.Next
		mark := false
		//删除重复的链表后的集合
		for q != nil && p.Val == q.Val {
			p.Next = q.Next
			q = q.Next
			mark = true
		}
		//判断当前p是否需要删除,若不需要删除则入栈
		if !mark {
			pre.Next = p
			p.Next = nil
			pre = pre.Next
		}
		p = q
	}
	head = result.Next
	return head
}
