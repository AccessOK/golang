package algorithms

import (
	"fmt"
	"sort"
)

func removeDuplicates(nums []int) int {
	//给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
	//考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
	//更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
	//返回 k 。
	//1 <= nums.length <= 3 * 10^4
	//-10^4 <= nums[i] <= 10^4
	//nums 已按 非严格递增 排列
	var k = 0
	var copy1 []int = make([]int, len(nums))
	var copy2 []int = make([]int, len(nums))
	copy(copy1, nums)
	sort.Ints(copy1)
	fmt.Println(copy1)
	fmt.Println(nums)
	var mark []int = make([]int, 10000)
	for i := 1; i < len(nums); i++ {
		if copy1[i] == copy1[i-1] {
			mark[copy1[i]]++
		}
	}
	fmt.Println(mark)
	var result []int = make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if mark[nums[i]] == 0 {
			result = append(result, nums[i])
		}
	}
	fmt.Println(result)
	nums = result
	return k
}
