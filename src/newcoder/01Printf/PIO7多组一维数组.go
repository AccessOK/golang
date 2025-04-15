package main

import "fmt"

func mainpio7() {
	t := 0
	fmt.Scanf("%d", &t)
	n := 0
	result := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		tmp := 0
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &nums[j])
			tmp += nums[j]
		}
		result[i] = tmp
	}
	for i := 0; i < t; i++ {
		fmt.Println(result[i])
	}
}
