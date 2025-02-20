
package main
import "fmt"
func main(){
	//内置函数 buitin下的包，不需要导入包可以直接使用
	fmt.Println(len("asda"))
	//new 函数 分配内存，返回分配的内存地址的指针
	ptr := new(int)
	fmt.Println(ptr)
	*ptr = 23
	fmt.Println(*ptr)
	
}