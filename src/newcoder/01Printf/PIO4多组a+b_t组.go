package main

import "fmt"

func mainpio4() {
	t := 0
	fmt.Scanf("%d", &t)
	a := 0
	b := 0
	for t > 0 {
		n, err := fmt.Scanf("%d %d", &a, &b)
		t--
		if n == 0 || err != nil {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}
