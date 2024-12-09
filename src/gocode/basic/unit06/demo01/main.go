package main
import ("fmt"
		"errors"
)

func main(){
	//错误处理+异常捕获机制
	//recover允许程序管理恐慌过程中的go程。
	//在defer的函数中，执行recover调用会取回传至panic调用的错误值，恢复正常执行。
	err := test()
	if err  != nil {
		fmt.Println(err)
		//终止错误程序
		panic(err)
	}
	fmt.Println("函数执行成功！")
}

func test()(err error){
	//利用defer+recover来捕获错误：defer + 匿名函数调用
	// defer func(){
	// 	//调用recover内置函数，可以捕获错误
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println("错误已经捕获！")
	// 	}
	// }()
	num1 := 10
	num2 := 0
	if num2 == 0 {
		//抛出自定义异常
		return errors.New("除数不能为0")
	}else{
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}