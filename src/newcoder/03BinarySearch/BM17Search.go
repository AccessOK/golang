package main

func Search(nums []int, target int) int {
	// write code here
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2
	for nums[mid] != target && right != left {
		if nums[mid] < target {
			//下一次二分的起始位置
			left = mid + 1
			//二分的位置相邻则没有可以再二分的位置了
			mid = (left + right) / 2
		} else if nums[mid] > target {
			right = mid
			mid = (left + right) / 2
		} else {
			break
		}
	}
	if nums[mid] == target {
		return mid
	} else {
		return -1
	}
}
