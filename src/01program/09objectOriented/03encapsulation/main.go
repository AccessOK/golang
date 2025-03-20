package main
import "fmt"
//外部包不能访问该结构体
//封装
type person struct{
	Name string
	age int  // 其他包不能访问age
}
//定义工厂模式函数，相当于构造器
func NewPerson(name string) * person{
	return &person{
		Name : name ,
	}
}
//定义set和get方法，对age字段进行封装
func (p *person) SetAge(age int){
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}
//定义set和get方法，对age字段进行封装
func (p *person) GetAge(age int){
	return p.age
}
