package algorithms

func removeDuplicates(nums []int) int {
	//给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
	//考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
	//更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
	//返回 k 。
	//1 <= nums.length <= 3 * 10^4
	//-10^4 <= nums[i] <= 10^4
	//nums 已按 非严格递增 排列

	mark := 99999
	//标记所有重复数据
	for i := 0; i < len(nums); i++ {
		if nums[i] != mark {
			for j := i + 1; j < len(nums); j++ {
				if i != j {
					if nums[j] != mark && nums[j] == nums[i] {
						nums[j] = mark
					}
				}
			}
		}
	}
	//统计非重复数据个数
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != mark {
			k++
		}
	}
	//迁移非重复数据
	for i := 0; i < len(nums); i++ {
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
	return k
}
