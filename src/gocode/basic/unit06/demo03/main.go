package main
import "fmt"
//可变参数 可以传任意多个数量的int类型的数据
func test(args...int){
	//函数内部处理可变参数的时候，可将可变参数当作切片处理
	for i := 0; i < len(args) ; i++ {
		fmt.Printf("aaa%d",args[i])
	}
}
func main(){
	test()
	fmt.Println("-----------")
	test(2)
	fmt.Println("-----------")
	test(4,12,32)
}