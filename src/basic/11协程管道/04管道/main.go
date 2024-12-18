package main
import (
	"fmt"
)
//管道 channel 引用类型，本质上是队列，先进先出
func main(){
	//声明管道
	var intChan chan int
	//通过make初始化：管道可以存放三个int类型的数据
	intChan = make(chan int,3)
	//证明管道是引用类型
	fmt.Printf("intChan = %v",intChan)
	//向管道存放数据
	intChan <- 10
	num := 20
	intChan <- num
	//输出管道的长度,管道的实际长度不能超出管道的容量
	fmt.Printf("管道的实际长度：%v,管道的容量：%v\n",len(intChan),cap(intChan))
	//在管道中读数
	num1 := <- intChan
	num2 := <- intChan
	// num3 := <- intChan 管道数据如果全部取出之后再取就会报错
	fmt.Println("num1 = ",num1)
	fmt.Println("num2 = ",num2)
	// fmt.Println("num3 = ",num3)
	intChan <- 99
	intChan <- 100
	intChan <- 101 
	//管道关闭，管道关闭后只出不进
	close(intChan)
	num4 := <- intChan
	fmt.Println("num4 = ",num4)
	// intChan <- 100 
	//管道遍历,管道没有索引，不能使用for i:=0 ,只能使用for range
	//在遍历前，如果没有关闭管带，就会出现deadlock错误
	for value := range intChan {
		fmt.Println("value = ", value)
	}

}