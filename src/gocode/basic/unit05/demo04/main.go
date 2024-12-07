package main
import "fmt"
//使用指针通过函数修改函数外的值
func test(num *int){
	*num = 20
}
func test1(num int,testfun func (*int))  {
	fmt.Println(num)
	fmt.Printf("fun type = %T",testfun)
}
//自定义函数类型
type myfunc func (*int)
func test2(num int,testfun myfunc)  {
	fmt.Println(num)
	fmt.Printf("fun type = %T\n",testfun)
}
//对函数的返回值命名，就不用固定return的顺序
func test3(num1 int,num2 int) (sum int,sub int){
	sum = num1 + num2
	sub = num1 - num2
	return
}
func main(){
	var num int = 10
	test(&num)
	fmt.Println(num)
	//函数也是数据类型，可以赋值给一个变量
	fun := test
	fmt.Printf("fun type = %T\n",fun)
	//函数可以作为一个参数传递
	test1(99,test)
	//自定义数据类型:(相当于起别名)
	type myInt int
	var num1 myInt = 30
	fmt.Println(num1)
	var num2 int = 20
	fmt.Println(num2)
	// num1=num2 报错，数据类型不匹配
	fmt.Println(num1)
	//自定义函数类型
	test2(99,test)
	//命名函数返回值
	sum,sub := test3(88,12)
	fmt.Println(sum)
	fmt.Println(sub)
}