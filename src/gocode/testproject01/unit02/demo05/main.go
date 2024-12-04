package main
import (
	"fmt"
	"strconv"
)

func main(){
	// string 与 与基本数据类型转换
	var n1 int =19
	var n2 float32 =3.14
	var n3 bool = false
	var n4 byte = 'a'
	var s1 string = fmt.Sprintf("%d",n1)
	fmt.Printf("s1的数据类型 = %T，s1的值 = %q\n",s1,s1)
	var s2 string = fmt.Sprintf("%f",n2)
	fmt.Printf("s2的数据类型 = %T，s2的值 = %q\n",s2,s2)
	var s3 string = fmt.Sprintf("%t",n3)
	fmt.Printf("s3的数据类型 = %T，s3的值 = %q\n",s3,s3)
	var s4 string = fmt.Sprintf("%c",n4)
	fmt.Printf("s4的数据类型 = %T，s4的值 = %q\n",s4,s4)

	// string 转换为 基本数据类型
	// 转换失败则返回默认值
	var str1 string = "true"
	var b bool
	//PareseBool 有两个返回值：(value bool,err error)
	//err使用_忽略
	b,_ =strconv.ParseBool(str1)
	fmt.Printf("b的类型%T，b的值%v",b,b)
}