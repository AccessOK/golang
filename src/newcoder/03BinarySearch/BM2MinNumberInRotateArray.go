package main

func minNumberInRotateArray(nums []int) int {
	// write code here
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	mid := (left + right) / 2
	for left < right {
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[left] {
			right = mid
		} else {
			right = right - 1
		}
		mid = (left + right) / 2
	}
	return nums[mid]
}
