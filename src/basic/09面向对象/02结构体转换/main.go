package main
import "fmt"
//定义结构体
type Teacher struct{
	//变量名首字母大写，外界可访问
	Name string
	Age int
	School string
}
//结构体的绑定方法
func (teacher Teacher) teach(){
	teacher.Name = "teach() change"
	fmt.Println("这是Teacher的teach()方法")
	fmt.Println(teacher)
}
//结构体的绑定方法
func (teacher *Teacher) teacha(){
	//* 可以省略，底层会自动添加
	(*teacher).Name = "teach() change"
	fmt.Println("这是Teacher的teach()方法")
	fmt.Println(teacher)
}
func (teacher *Teacher) String() string{
	str := fmt.Sprintf("Name = %v,Age = %v",teacher.Name,teacher.Age)
	return str
}
func main(){
	var teacher1 Teacher
	teacher1.Name = "sdas"
	teacher1.Age = 99
	//值传递，副本传递，不影响当前函数的内存值
	teacher1.teach()
	fmt.Println(teacher1)
	//指针传递
	(&teacher1).teacha()
	fmt.Println(teacher1)
	fmt.Println(&teacher1)
}