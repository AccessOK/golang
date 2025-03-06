package algorithms

import "fmt"

func removeDuplicatesPlus(nums []int) int {
	//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
	//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
	//说明：
	//为什么返回数值是整数，但输出的答案是数组呢？
	//请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
	// 在函数里修改输入数组对于调用者是可见的。
	// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。

	// 标记数量超出2位的重复数据
	mark := 99999
	for i := 0; i < len(nums); i++ {
		number := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				number++
				if number >= 2 {
					nums[j] = mark
				}
			}
		}
	}
	fmt.Println(nums)
	//标记k的数量
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != mark {
			k++
		}
	}
	fmt.Println(k)
	//迁移数据
	for i := 1; i < len(nums); i++ {
		if nums[i] == mark {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != mark {
					nums[i] = nums[j]
					nums[j] = mark
					break
				}
			}
		}

	}
	fmt.Println(nums)
	return k
}
