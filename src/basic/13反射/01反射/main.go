package main
import (
	"fmt"
	"reflect"
)
//对结构体反射
type student struct{
	Name string
	Age int
}

//反射可以在运行时动态获取变量的各种信息，比如变量的类型，类别
//通过反射，可以修改变量的值，调用关联的方法
func testReflect(i interface{}){
	//调用TypeOf函数，返回reflect.Type类型数据
	reType := reflect.TypeOf(i)
	fmt.Println("reType:",reType)
	fmt.Printf("reType的具体类型是：%T",reType)
	//调用ValueOf函数，返回reflect.Value类型数据
	reValue := reflect.ValueOf(i)
	fmt.Println("reValue:",reValue)
	fmt.Printf("reValue的具体类型是：%T\n",reValue)
	//如果真想获取reValue的值，需要调用Int()方法
	num := 80 + reValue.Int()
	fmt.Println(num)
}
//反射结构体
func testReflectStruct(i interface{}){
	//调用TypeOf函数，返回reflect.Type类型数据
	reType := reflect.TypeOf(i)
	fmt.Println("reType:",reType)
	fmt.Printf("reType的具体类型是：%T",reType)
	//调用ValueOf函数，返回reflect.Value类型数据
	reValue := reflect.ValueOf(i)
	fmt.Println("reValue:",reValue)
	fmt.Printf("reValue的具体类型是：%T\n",reValue)
	//reValue转成空接口
	i2 := reValue.Interface()
	n,flag := i2.(student)
	if flag == true{
		fmt.Printf("Name = %v ,Age = %v\n",n.Name,n.Age)
	}
	//查看类别和类型
	k1 := reType.Kind()
	fmt.Println(k1)
	k2 := reValue.Kind()
	fmt.Println(k2)
	fmt.Printf("%T\n",n)
}
//修改反射出来的基本类型的值
func changeReflect(i interface{}){
	reValue := reflect.ValueOf(i)
	reValue.Elem().SetInt(88)
}
func main(){
	//对基本数据类型进行反射
	var num int = 100
	testReflect(num)

	//反射结构体
	student1 := student{"lilo",76}
	testReflectStruct(student1)

	//修改值
	var num1 int = 99
	changeReflect(&num1)
	fmt.Println(num1)
}