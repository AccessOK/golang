package main
import "fmt"
//内存分析
func exchange(a int,b int)(int,int){
	var t int = b
	b = a
	a = t
	return a,b
}
func main(){
	var num1 int = 5
	var num2 int = 9
	exchange(num1,num2)
	fmt.Printf("num1=%d,num2=%d",num1,num2)
	num1,num2=exchange(num1,num2)
	fmt.Printf("num1=%d,num2=%d",num1,num2)
}