package algorithms

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	majorityElement(nums)
}

func TestMerge(t *testing.T) {
	var num1 = []int{0, 1, 2, 3, 0, 0, 0}
	var num2 = []int{2, 4, 5}
	merge(num1, 4, num2, 3)
}
func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	//nums := []int{1, 2, 3}
	rotate(nums, 5)
}
func TestRemoveElement(t *testing.T) {
	var nums []int = []int{0, 1, 2, 2, 3, 0, 4, 2} // 输入数组
	var val = 2                                    // 要移除的值
	removeElement(nums, val)                       // 调用实现
}
func TestRemoveDuplicatesPlus(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3} // 输入数组
	removeDuplicatesPlus(nums)
}
func TestRemoveDuplicates(t *testing.T) {
	var nums []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4} // 输入数组
	removeDuplicates(nums)
}

func TestMaxProfit(t *testing.T) {

}
