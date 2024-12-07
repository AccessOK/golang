package main
import "fmt"
//匿名函数，只是用意一次的函数可以定义为匿名函数

//将匿名函数赋给全局变量
var func01 = func (num1 int,num2 int) int {
	return num1 * num2
}

func main(){
	result1 := func (num1 int,num2 int) int {
		return num1 + num2
	}(10,20)
	fmt.Println(result1)
	//将匿名函数赋给一个变量，这个变量实际就是函数类型的变量
	sub := func (num1 int,num2 int) int {
		return num1 - num2
	}
	result2 := sub(10,9)
	fmt.Println(result2)
	result03 := func01(8,2)
	fmt.Println(result03)
}