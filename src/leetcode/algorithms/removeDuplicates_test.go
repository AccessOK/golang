package algorithms

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var nums []int = []int{0, 1, 2, 2, 3, 0, 4, 2} // 输入数组
	removeDuplicates(nums)
}
