package main
import "fmt"
func exchange(a int,b int)(int,int){
	var t int = b
	b = a
	a = t
	return a,b
}
func main(){
	var num1 int = 5
	var num2 int = 9
	exchanpackage main
	import "fmt"
	//形参定义 类型在参数名后面
	//函数名首字母大写可以被本报文件和其他文件使用（public）
	func sum(a int,b int)(int){//如果返回类型就一个，可以省略
		return a + b
	}
	func cal(a int,b int)(int,int){
		return a-b,b-a
	}
	func nothing(){
		fmt.Println("nothing")
	}
	func main(){
		fmt.Println(sum(9,1))
		fmt.Println(cal(9,3))
		nothing()
	}ge(num1,num2)
	fmt.Printf("num1=%d,num2=%d",num1)
}