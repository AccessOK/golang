package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	num := make([]int, m+n)
	markM := 0
	markN := 0
	for i := 0; i < m+n; i++ {
		if nums1[markM] < nums2[markN] {
			num[i] = nums1[markM]
			markM++
		} else if markN < n {
			num[i] = nums2[markN]
			markN++
		}
	}
	fmt.Println(markM, "\n", markN, "\n", num, "\n")
}
func main() {
	fmt.Println("this is main")
	//合并两个有序数组
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := len(nums1)
	nums2 := []int{1, 3, 4, 5, 6, 8, 9, 10}
	n := len(nums2)
	merge(nums1, m, nums2, n)
}
