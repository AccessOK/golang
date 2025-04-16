package BinarySearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{}
	nums = append(nums, 1, 0, 1, 1, 1)
	//result := Search(nums, 8)
	result := minNumberInRotateArray(nums)
	fmt.Println(result)
}
