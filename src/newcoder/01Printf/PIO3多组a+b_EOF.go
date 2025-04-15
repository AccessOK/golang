package main

import (
	"fmt"
)

func mainpio3() {
	a := 0
	b := 0
	for {
		n, err := fmt.Scanf("%d %d", &a, &b)
		if n == 0 || err != nil {
			break
		} else {
			fmt.Println(a + b)
		}
	}
}
