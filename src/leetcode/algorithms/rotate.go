package algorithms

import "fmt"

func rotate(nums []int, k int) {
	//给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
	//创建数组暂存k个值
	//若移动位数超过数组本身长度，则相当于循环轮转
	if k >= len(nums) {
		k = k % len(nums)
	}
	//若向右移动次数小于len/2，则向右移东更快
	if k <= len(nums)/2 {
		tmp := make([]int, k)
		//暂存最右侧的k个数
		tmpMarkIn := 0
		for i := len(nums) - 1; i > len(nums)-1-k; i-- {
			tmp[tmpMarkIn] = nums[i]
			tmpMarkIn++
		}
		//从最右侧开始遍历移动
		tmpMarkOut := 0
		for i := len(nums) - 1; i >= 0; i-- {
			if i > k-1 {
				nums[i] = nums[i-k]
			} else {
				nums[i] = tmp[tmpMarkOut]
				tmpMarkOut++
			}
		}
	} else {
		//若向右移动次数大于len/2，则向左移东更快
		k = len(nums) - k
		//暂存向左移动k位
		tmp := make([]int, k)
		//暂存最左侧的k个数
		tmpMarkIn := 0
		for i := 0; i < k; i++ {
			tmp[tmpMarkIn] = nums[i]
			tmpMarkIn++
		}
		//从最左侧开始遍历移动
		tmpMarkOut := 0
		for i := 0; i < len(nums); i++ {
			if i < len(nums)-k {
				nums[i] = nums[i+k]
			} else {
				nums[i] = tmp[tmpMarkOut]
				tmpMarkOut++
			}
		}
	}
	fmt.Println(nums)
}
