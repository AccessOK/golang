package main
import "fmt"
//引入包的时候会先导入该包的全局变量和init函数
//1.定义全局变量
var num int = test()
func test() (int){
	fmt.Println("test")
	return 1
}
//2.init 初始化函数，init在main前执行
func init(){
	fmt.Println("init")
}
//3.main函数调用
func main(){
	fmt.Println(num)
	fmt.Println("main")
}