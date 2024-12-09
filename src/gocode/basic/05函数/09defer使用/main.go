package main
import "fmt"
//defer 释放资源
func main(){
	sum := add(2,4)
	fmt.Println("sum=",sum)
}
func add(num1 int, num2 int) (sum int){
	//defer：不会立即执行defer后的语句，
	//而是将defer后的语句压入一个栈中，然后执行函数后面的的语句，defer 先入后出
	//defer 会将相关值同时拷贝入栈中，不会随着函数后面变化而变化
	//应用场景：关闭某个使用资源，使用defer执行，函数执行完之后会自动关闭和释放资源
	defer fmt.Println("num1=",num1)
	defer fmt.Println("num2=",num2)
	num1 += 2
	num2 += 3
	sum = num1 + num2
	fmt.Println("add func sum = ",sum)
	return
}