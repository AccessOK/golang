package meeting

import "fmt"

func main1() {
	own := []int{}
	for {
		tmp := 0
		fmt.Scanf("%d", &tmp)
		own = append(own, tmp)
	}
	all := []int{}
	for {
		tmp := 0
		fmt.Scanf("%d", &tmp)
		all = append(all, tmp)
	}

}
