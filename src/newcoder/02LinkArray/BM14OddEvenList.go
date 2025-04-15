package main

func OddEvenList(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	count := 1
	prevHalf := head
	prevHead := prevHalf
	lastHalf := head.Next
	lastHead := lastHalf
	for head != nil && head.Next != nil {
		//当前为奇数节点
		if count%2 == 1 {
			prevHead.Next = head.Next.Next
			if prevHead.Next != nil {
				prevHead = prevHead.Next
			}
			head = lastHead
		} else {
			lastHead.Next = head.Next.Next
			if lastHead.Next != nil {
				lastHead = lastHead.Next
			}
			head = prevHead
		}
		count++
	}
	prevHead.Next = lastHalf
	return prevHalf
}
