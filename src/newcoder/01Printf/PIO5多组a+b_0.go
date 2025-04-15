package main

import "fmt"

func mainpio5() {
	a := 0
	b := 0
	for {
		fmt.Scanf("%d %d", &a, &b)
		if a == 0 && b == 0 {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}
