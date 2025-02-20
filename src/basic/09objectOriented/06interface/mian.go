package main
import "fmt"
//定义接口,接口默认是一个指针，空接口可以接任何类型
//接口中可以定义一组方法，但不需要实现，不需要方法体，并且接口中不能包含任何变量
//到某个自定义类型要使用的时候，在根据具体情况把这些方法实现出来
//实现接口要实现所有的方法才叫实现

type SayHello interface{
	//声明没有实现的方法
	sayHello()
}

//实现接口
type Chinese struct{
	name string
}

func (person Chinese) sayHello() {
	fmt.Println("你好")
}

type American struct{
	name string
}

func (person American) sayHello() {
	fmt.Println("Hello")
}
//多态实现	
//多态通过接口实现，可以通过统一的接口来调用不同的实现
func greet(s SayHello){
	s.sayHello()
}

//只要是自定义数据类型都可以实现接口，不仅仅是结构体类型
type integer int
func (i integer) sayHello(){
	fmt.Println("hello ",i)
}

//一个自定义类型可以实现多个接口
type SayWelcome interface{
	sayWelcome()
}
func (person Chinese) sayWelcome() {
	fmt.Println("欢迎")
}

func main(){
	//创建中国人
	c := Chinese{}
	//创建美国人
	a := American{}
	greet(a)
	greet(c)
	//接口本身不能创建实例，但是可以只想一个实现了该接口的自定义类型变量
	// var s SayHello
	// s.sayHello()
	var s SayHello = c 
	s.sayHello()
	var i integer = 10
	var s1 SayHello = i
	s1.sayHello()
	var s2 SayHello = c 
	var s3 SayWelcome = c 
	s2.sayHello()
	s3.sayWelcome()
	//多态数组
	var person [3]SayHello
	person[0] = American{"acwe"}
	person[1] = American{"wewq1"}
	person[2] = Chinese{"自拍摄的"}
	fmt.Println(person)
}