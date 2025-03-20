package main
import (
	"fmt"
	"unsafe"
)
func main(){
	fmt.Println("--------------------------整形------------------------------------")
	//超出字节范围
	//var num int8 = 230
	var num int16 = 230
	fmt.Println(num)

	//var num2 uint8 = -12
	var num2 uint8 = 12
	fmt.Println(num2)
	var num3 = 12
	fmt.Printf("num3的默认类型判断为%T",num3)
	fmt.Println(unsafe.Sizeof(num3))

	fmt.Println("--------------------------浮点类型------------------------------------")
	//定义浮点类型数据
	var fnum1 float32 = 3.14
	fmt.Println(fnum1)
	//定义负浮点类型数据
	var fnum2 float32 = -3.14
	fmt.Println(fnum2)
	//浮点数据可以使用科学计数法表示，E大写小写都行
	var fnum3 float32 = 314E-2
	fmt.Println(fnum3)
	var fnum4 float32 = 314E+2
	fmt.Println(fnum4)
	var fnum5 float32 = 314e+2
	fmt.Println(fnum5)
	var fnum6 float64 = 314E+2
	fmt.Println(fnum6)
	//浮点数会有精度损失，建议使用：float64
	var fnum7 float32 = 256.0000000916
	fmt.Println(fnum7)
	var fnum8 float64 =  256.0000000916
	fmt.Println(fnum8)
	//golang中默认的浮点类型为：float64
	var fnum9 = 3.17
	fmt.Printf("fnum9对应的默认类型: %T",fnum9)
}