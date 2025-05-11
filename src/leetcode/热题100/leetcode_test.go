package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(f *testing.F) {
	nums := []int{2, 7, 11, 15}
	target := 9
	total := twoSum(nums, target)
	fmt.Println(total)
}
