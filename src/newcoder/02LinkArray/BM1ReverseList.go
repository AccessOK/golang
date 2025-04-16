package LinkArray

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @return ListNode类
 */

func ReverseList(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	prev := head
	current := head.Next
	head = head.Next
	prev.Next = nil
	for head != nil {
		head = head.Next
		current.Next = prev
		prev = current
		current = head
	}
	head = prev
	return head
}
