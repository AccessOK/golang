package algorithms

import "fmt"

func rotate(nums []int, k int) {
	//给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

	//创建数组暂存k个值
	tmp := make([]int, k)
	tmpMarkIn := 0
	tmpMarkOut := 0
	if len(nums) <= k {
		k = k % len(nums)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i > len(nums)-1-k {
			tmp[tmpMarkIn] = nums[i]
			tmpMarkIn++
			nums[i] = nums[i-k]
		} else if i < k {
			nums[i] = tmp[tmpMarkOut]
			tmpMarkOut++
		} else {
			nums[i] = nums[i-k]
		}
	}
	fmt.Println(nums)
}
