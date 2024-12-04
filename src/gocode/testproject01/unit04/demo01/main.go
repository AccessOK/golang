package main
import "fmt"
func main(){
	// if 后面可以并列的加入变量的定义
	if count := 21.1;count <20 {
		fmt.Println(count)
	}else if count == 20{
		fmt.Println("sorry!")
	}else {
		fmt.Println("last!")
	}

	// switch 
	// case里面不需要break，default也不是必须的，case 后可以跟多个值
	var score int = 17
	switch score/10 {
		case 10,9:
			fmt.Println("A")
		case 8:
			fmt.Println("C")
			//fallthrough 可以穿透到下一层case
			fallthrough
		case 7:
			fmt.Println("D")
		default:
			fmt.Println("E")
	}

	//for
	var sum int =0
	for i := 1 ; i <= 5 ; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 灵活for
	sum=0
	i := 1
	for  ; i <= 5 ; {
		sum += i
		i++
	}
	fmt.Println(sum)

	//死循环
	// sum=0
	// for {
	// 	sum += i
	// 	i++
	// 	fmt.Println(sum)
	// }	sum=0
	// for;; {
	// 	sum += i
	// 	i++
	// 	fmt.Println(sum)
	// }

	// for range
	var str string = "hello world  你哈"
	// for i := 0 ; i < len(str) ; i++{
	// 	fmt.Printf("%c",str[i])
	// }
	for i,value := range str {
		fmt.Printf("%d %c,",i,value)
	}

	//break 跳出一层循环，跳出到标签循环
	lable1:
	for i := 0 ; i < 5 ; i++ {
		for j:=0 ; j< 4 ; j++ {
			fmt.Println(j)
			if i==2 && j==2 {
				break lable1
			}
		}
		fmt.Println(i)
	}

	//continnue
	
}