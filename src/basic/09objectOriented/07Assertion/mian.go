package main

import "fmt"

// 断言，判断是否是该类型的变量
// 用来验证程序在运行时某个条件是否为真。如果条件为假，程序通常会抛出错误或中断执行。它主要用于调试阶段，帮助开发者发现潜在的逻辑错误或不符合预期的情况。
// 定义接口
type interfaceA interface {
	test()
}

// 定义结构体
type typeA struct {
}

// 实现接口
func (a typeA) test() {
	fmt.Println("test()")
}

// 定义typeA结构体特有的函数
func (a typeA) testA() {
	fmt.Println("testA()")
}

// 多态使用
func masd(i interfaceA) {
	i.test()
	fmt.Println("i test()")
	//断言：判断i是否可以转换成interfaceA
	// ii,flag := i.(typeA)
	// if flag ==true {
	// 	ii.testA()
	// }else{
	// 	fmt.Println("error")
	// }
	//断言简写
	// if ii,flag := i.(typeA);flag ==true {
	// 	ii.testA()
	// }else{
	// 	fmt.Println("error")
	// }
	switch i.(type) { // type属于go中的一个关键字，固定写法
	case typeA:
		fmt.Println("asd")
	}
}
func main() {
	var a interfaceA = typeA{}
	a.test()
	masd(a)
}
