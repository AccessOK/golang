package algorithms

import "testing"

func TestMerge(t *testing.T) {
	var num1 = []int{0, 1, 2, 3, 0, 0, 0}
	var num2 = []int{2, 4, 5}
	merge(num1, 4, num2, 3)
}
