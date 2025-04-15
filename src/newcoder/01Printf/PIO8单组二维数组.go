package main

import "fmt"

func mainpio8() {
	n := 0
	m := 0
	fmt.Scanf("%d %d", &n, &m)
	nums := [][]int{}
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < m; j++ {
			tmp := 0
			fmt.Scanf("%d", &tmp)
			row = append(row, tmp)
		}
		nums = append(nums, row)
	}
	sum := 0
	for _, row := range nums {
		for _, num := range row {
			sum += num
		}
	}
	fmt.Println(sum)
}
