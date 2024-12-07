package main
import "fmt"
//函数名为getSum 参数为空
//getSum 返回值是一个函数，该函数的参数是一个int类型的参数，返回值也是int
//闭包：返回的额匿名函数+匿名函数以外的变量
//闭包就是一个函数和与其相关的引用环境组合的一个整体
func getSum() func(int) int{
	var sum int = 0
	//返回一个匿名函数 
	return func (num int) int{
		sum = sum + num
		return sum
	}

}

func main(){
	f := getSum()
	var result1 = f(19)
	fmt.Println(result1)
	var result2 = f(120)
	fmt.Println(result2)
}