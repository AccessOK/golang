package main

import "fmt"

func mainpio6() {
	t := 0
	fmt.Scanf("%d", &t)
	a := make([]int, t)
	sum := 0
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &a[i])
		sum += a[i]
	}
	fmt.Println(sum)
}
