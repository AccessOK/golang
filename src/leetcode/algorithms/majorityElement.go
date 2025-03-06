package algorithms

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
	//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
	sort.Ints(nums)
	fmt.Println(nums)
	k := 0
	mark := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			k++
			if k >= len(nums)/2 {
				mark = i
			}
		} else {
			k = 0
		}
	}
	return nums[mark]
}
