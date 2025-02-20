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

func (s student) Cprint(){
	fmt.Println("Cprint()")
	fmt.Println("Name = ",s.Name)
}
func (s student) AetSum(n1,n2 int) int{
	fmt.Println("GetSum()")
	return n1 + n2
}
func (s student) BSet(name string, age int){
	fmt.Println("BSet()")
	s.Name = name
	s.Age = age
}
//定义函数操作结构体进行反射操作
func TestStudentStruct(a interface{}){
	//将a转成reflect.Value类型
	val := reflect.ValueOf(a)
	fmt.Println("val = ",val)
	//通过reflect.Value类型操作结构体的内部的字段
	n1 := val.NumField()
	fmt.Println(n1)
	for i := 0 ; i < n1 ; i++ {
		fmt.Printf("第%d个字段是：%v\n",i,val.Field(i))
	}
	//通过reflect.value类型操作结构体内部的方法
	n2 := val.NumMethod()
	fmt.Println(n2)

	//调用Cprint()方法：
	//调用方法，方法的首字母必须大写才能有对应的反射的访问权限
	//方法的顺序按照ASCII的顺序排列
	val.Method(2).Call(nil)

	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(21))
	result := val.Method(0).Call(params)
	fmt.Println("result=",result[0])
}
// 修改结构体的值
func fixStudentStruct(i interface{}){
	val := reflect.ValueOf(i)
	fmt.Println(val)
	val.Elem().Field(0).SetString("wj")
}
func main(){
	s := student{"asd",22}
	TestStudentStruct(s)
	fmt.Println(s)
	fixStudentStruct(&s)
	fmt.Println(s)
}