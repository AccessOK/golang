package main
import "fmt"
//定义结构体
type Teacher struct{
	//变量名首字母大写，外界可访问
	Name string
	Age int
	School string
}
type Student struct{
	//变量名首字母大写，外界可访问
	Name string
	Age int
	School string
}
//定义别名
type Person Student
func main(){
	//结构体做对象
	//1.先声明在赋值
	var teacher1 Teacher 
	fmt.Println(teacher1)
	teacher1.Name = "wjas"
	teacher1.Age = 23
	teacher1.School = "sad"
	fmt.Println(teacher1)
	//申明的时候赋值
	var teacher2 Teacher  = Teacher{"sd",32,"dsaseqw"}
	fmt.Println(teacher2)
	//teacher3是指向内存的地址
	var teacher3 *Teacher = new(Teacher)
	(*teacher3).Name = "asdsd"
	(*teacher3).Age = 99
	(*teacher3).School = "asdweq"
	fmt.Println(*teacher3)
	//teacher4是指向内存的地址并直接赋值
	var teacher4 *Teacher = &Teacher{"we",123,"343"}
	fmt.Println(teacher4)
	//结构体转换
	//字段必须完全相同
	var student1 Student 
	//student1 = teacher4 不能直接使用等号转换
	student1 = Student(teacher1)
	fmt.Println(student1)
	//使用别名也需要强制转换
	var person1 Person 
	person1 = Person(student1)
	fmt.Println(person1)
}
