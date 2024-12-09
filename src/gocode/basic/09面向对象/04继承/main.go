package main
import "fmt"
//继承
//定义动物的结构体：抽出共有的属性和方法
type Animal struct{
	Age int
	Weight float32
}
func (an *Animal) Shout(){
	fmt.Println("喊叫")
}
func (an *Animal) ShowInfo(){
	fmt.Printf("Age = %v, Weight = %v\n",an.Age,an.Weight)
}
//定义子类
type Cat struct{
	//为了服用型，体现继承思维，加入匿名结构体
	Animal
	//就近原则测试：编译器就近使用子类的属性和方法然后使用父类的属性和方法
	Age int
}
func (c *Cat) ShowInfo(){
	fmt.Printf("``````````````Age = %v, Weight = %v\n",c.Age,c.Weight)
}
//多继承
//支持多继承，但是不建议使用，防止使用时属性混乱
type CuteCat struct{
	Animal
	Cat
}

//对Cat 绑定特有方法
func (cat *Cat) scratch(){
	fmt.Println("挠人")
}
func main(){
	//创建cat结构体
	cat1 := &Cat{}
	cat1.Animal.Age = 3
	cat1.Animal.Weight = 10.3
	cat1.Animal.Shout()
	cat1.Animal.ShowInfo()
	cat1.scratch()
	//匿名结构体字段访问可以简化
	//先找子类的属性和方法，然后自动查找父类的属性和方法
	//编译请会采用就近访问原则
	cat2 := &Cat{}
	cat2.Age = 3
	cat2.Weight = 10.3
	cat2.Shout()
	cat2.ShowInfo()
	cat2.scratch()
	//就近原则测试
	cat3 := &Cat{}
	cat3.Age = 99
	cat3.Weight = 10.6//就近原则
	cat3.Shout()
	cat3.ShowInfo()//就近原则
	cat3.Animal.Age = 100
	cat3.Animal.ShowInfo()
	cat3.scratch()
	//多继承定义
	cuteCat := &CuteCat{Animal{10,29.1},Cat{Animal{1,23.12},99}}
	fmt.Println(*cuteCat)
	//嵌入结构体有相同的字段名或者方法名，在访问时，需要通过匿名结构体类型来区分
	//结构体的匿名字段可以时基本数据类型和指针
}