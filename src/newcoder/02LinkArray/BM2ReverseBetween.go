package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here
	count := 1
	result := head
	var startPrev *ListNode = nil
	var endLast *ListNode = nil
	var startLast *ListNode = nil
	var endPrev *ListNode = nil
	for head != nil {
		if count == m-1 {
			startPrev = head
		}
		if count == n+1 {
			endLast = head
			endPrev.Next = endLast
		}
		//首个反转节点
		if count == m {
			//到达反转的位置
			endPrev = head
			prev := head
			head = head.Next
			last := head.Next
			//当前位置为反转后的最后一个位置
			head.Next = prev
			prev.Next = nil
			count++
			//开始向后便利
			for count <= n && head != nil {
				if count == n {
					startLast = head
					startPrev.Next = startLast
					head = last
					count++
					break
				}
				prev = head
				head = last
				last = last.Next
				head.Next = prev
				count++
			}
		} else {
			head = head.Next
			count++
		}
	}
	return result
}
