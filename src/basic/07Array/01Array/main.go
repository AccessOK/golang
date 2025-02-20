package main
import "fmt"
func main(){
	//数组
	var scores [5]int
	//键盘录入
	// for i := 0;i < len(scores) ; i++ {
	// 	fmt.Scanln(&scores[i])
	// }
	scores = [5]int{2,3,55,63,2}
	scores = [...]int{2,3,5,3,3}
	scores = [...]int{2:22,4:99}
	for i := 0;i < len(scores) ; i++ {
		fmt.Println(scores[i])
	}
	for j,value := range scores {
		fmt.Println(j,value)
	}
	//数组的长度属于类型的一部分
	//函数传递属于值传递，不能改变源变量的值
	//若想通过函数调用修改源变量的值，需要传递指针
}