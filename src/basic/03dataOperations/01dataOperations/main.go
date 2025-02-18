package main
import "fmt"
func main(){
	// +
	//1.正数，2.加操作，3.字符串拼接
	var n1 int = +10
	fmt.Println(n1)
	var n2 int = 3
	fmt.Println(n1+n2)
	var s1 string = "abc" + "def"
	fmt.Println(s1)

	// /
	fmt.Println(10/3)
	fmt.Println(10.0/3)

	// %
	fmt.Println(10%3)

	// ++ 只能单独使用，不能参与到运算中去，只能在变量的后面，不能在变量的前面
	var a int = 11
	a++
	fmt.Println(a)

	// = += /= *=
	var num int = 10
	num = num +20
	fmt.Println(num)
	num += 20
	fmt.Println(num)

	// == !-= < > <= >=

	// && || !  短路与，短路或，提高运算效率
}