package LinkArray

import (
	"sort"
)

func sortInList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	a := []int{}
	for p != nil {
		a = append(a, p.Val)
		p = p.Next
	}
	p = head
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		p.Val = a[i]
		p = p.Next
	}
	return head
	// 选择排序，超时
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//var resultNode = new(ListNode)
	//result := resultNode
	////第一层循环
	//first := head
	////记录最值点
	//for first != nil {
	//	second := first
	//	minPrev := second
	//	minNode := second
	//	//找出当前队列中的最小值
	//	for second != nil && second.Next != nil {
	//		if minNode.Val > second.Next.Val {
	//			minPrev = second
	//			minNode = second.Next
	//		}
	//		second = second.Next
	//	}
	//	//找到最小值
	//	if minPrev != minNode {
	//		minPrev.Next = minNode.Next
	//	} else {
	//		first = first.Next
	//	}
	//	minNode.Next = nil
	//	resultNode.Next = minNode
	//	resultNode = resultNode.Next
	//}
	//head = result.Next
	//return head
}
