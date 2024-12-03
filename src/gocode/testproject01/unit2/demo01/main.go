package main

import "fmt"

// 全局变量
var global1 = 100
var global2 = "qwe"

//全局变量的一次性声明
var (
	global3 = 990
	global4 = "sdss"
)
func main() {

	fmt.Println(global1)
	fmt.Println(global2)
	fmt.Println(global3)
	fmt.Println(global4)

	fmt.Println("-------------------------------------")
	//变量的声明
	var age int
	//变量的赋值
	age = 10
	//变量的使用
	fmt.Println("age=", age)

	//变量不可以重复定义
	//变量类型必须匹配

	fmt.Println("-------------------------------------")
	//变量的四种使用方式

	//1.指定变量类型并且赋值
	var num1 int = 1
	fmt.Println(num1)
	//2.指定变量的类型，但不赋值，使用默认值
	var num2 int
	fmt.Println(num2)
	//3.不写变量类型，根据赋值进行判断
	var num3 = "tom"
	fmt.Println(num3)
	//4.省略var，注意 := 不能写为 =
	num4 := "1s"
	fmt.Println(num4)

	fmt.Println("-------------------------------------")
	// 一次性声明多个变量
	var num5,num6,num7 int
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)

	var num8,num9,num10 = 10,2,"hello"
	fmt.Println(num8)
	fmt.Println(num9)
	fmt.Println(num10)

	num11,num12 := 99,"hwoe"
	fmt.Println(num11)
	fmt.Println(num12)
}
