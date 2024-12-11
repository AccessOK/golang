package main
import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func writeData(intChan chan int){
	defer wg.Done()
	for i := 0 ;i < 10 ; i++ {
		intChan <- i
		fmt.Println("写入数据：",i)
		time.Sleep(time.Second)
	}
	close(intChan) //写完数据之后关闭管道
}
func readData(intChan chan int){
	defer wg.Done()
	for value := range intChan {
		fmt.Println("读数据：",value)
		time.Sleep(time.Second)
	}
}
func main(){
	//默认情况下管道是双向的，可读可写
	var intChan chan int
	intChan = make(chan int,10)
	wg.Add(2)
	go writeData(intChan)
	go readData(intChan)
	wg.Wait()
	fmt.Println("主线程结束")
	
	//管道可以什么读写权限
	var intChanW chan<- int //只写管道
	intChanW = make(chan int,3)
	intChanW<- 20
	// num5 := <- intChanW 权限报错
	// fmt.Println(num5)
	fmt.Println(intChanW)
	var intChanR <-chan int //只读管道
	if intChanR != nil {
		num6 := <- intChanR
		fmt.Println(num6)
	}
	// intChanR<- 30 权限报错

	//管道堵塞，管道只写入，不读取的时候会堵塞

}