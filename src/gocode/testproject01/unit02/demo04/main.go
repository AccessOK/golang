package main
import "fmt"

func main(){
	//类型转换
	var n1 int = 20
	fmt.Println(n1)
	// var n2 float32 = n1 转换失败
	var n2 float32=float32(n1)
	fmt.Println(n2)
	fmt.Printf("%T\n",n1)
	fmt.Printf("%T\n",n2)

	//将int64转为int8的时候，编译不会出错，但是数据会溢出
	var n3 int64 = 9999
	fmt.Println(n3)
	var n4 int8=int8(n3)
	fmt.Println(n4)
	fmt.Printf("%T\n",n3)
	fmt.Printf("%T\n",n4)

	//等号左右的数据类型一定要相同
	var n5 int32 = 12;
	fmt.Println(n5)
	// var n6 int64 = n5 + 30  编译报错
	var n6 int64 = int64(n5) + 30
	fmt.Println(n6)

	var n7 int64 = 12
	var n8 int8 = int8(n7) + 127 //0000 1100 + 0111 1111 = 1000 1011 ，1000 1011 取反码为0111 0100，-(反码+1)=-117
	fmt.Println(n8)
	// var n9 int8 = int8(n7) + 128 //编译失败,128 本身以及超出int8可以表述的范围了
	// fmt.Println(n9)
}