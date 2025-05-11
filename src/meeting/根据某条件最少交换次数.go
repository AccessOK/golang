package meeting

import "fmt"

func minChange() {
	nums := make([]int, 100)
	numbers := 0
	for {
		tmp := 0
		fmt.Scanf("%d", nums[numbers])
		nums = append(nums, tmp)
		numbers++
		if tmp > 100 || tmp < -100 {
			break
		}
	}
	k := 0
	fmt.Scanf("%d", &k)
	//
}
