package algorithms

import "testing"

func TestRemoveElement(t *testing.T) {
	var nums []int = []int{0, 1, 2, 2, 3, 0, 4, 2} // 输入数组
	var val = 2                                    // 要移除的值
	removeElement(nums, val)                       // 调用实现
}
