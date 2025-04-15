package main

import "fmt"

func mainpio1() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scanf("%d %d", &a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("Hello Nowcoder!\n")
		}
	}
}
