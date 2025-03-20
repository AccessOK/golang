package main

import "fmt"

func main() {
	var age int = 18
	//打印age的地址
	fmt.Println(&age)
	//定义一个指针变量，ptr对应的类型是指向int类型的指针，&age为指向的实际地址的值
	var ptr *int = &age
	//输出ptr指向地址的实际值
	fmt.Println(*ptr)

	//可以通过指针改变变量值
	*ptr = 20
	fmt.Println(*ptr)
	fmt.Println(age)
	//指针变量接受的一定是地址值
	// ptr = 20
	//指针变量的地址不可以不匹配
	//var ptr1 *float32 = age
	//fmt.Println(*ptr1)

}
