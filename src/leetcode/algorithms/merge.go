package algorithms

import "fmt"

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
// 为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。nums1 []int, m int, nums2 []int, n int
func merge(nums1 []int, m int, nums2 []int, n int) {
	var m_mark = m
	var n_mark = n
	for i := len(nums1); i > 0; i-- {
		fmt.Printf("%d\n", m_mark)
		fmt.Printf("%d\n", n_mark)
		if n_mark > 0 && m_mark > 0 {
			if nums1[m_mark-1] > nums2[n_mark-1] && m_mark > 0 {
				nums1[i-1] = nums1[m_mark-1]
				m_mark--
			} else if n_mark > 0 {
				nums1[i-1] = nums2[n_mark-1]
				n_mark--
			}
		} else if n_mark > 0 {
			nums1[i-1] = nums2[n_mark-1]
			n_mark--
		} else {
			nums1[i-1] = nums1[m_mark-1]
			m_mark--
		}
		fmt.Println(nums1)
		fmt.Println(nums2)
	}
	fmt.Println(nums1)
}
