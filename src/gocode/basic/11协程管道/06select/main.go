package main
import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup


func main(){
	//select 多管道选择问题，也可以叫多路复用，可以从多个管道中随机公平的选择一个来执行
	//定义两个管道
	intChan := make(chan int,5)
	stringChan := make(chan string,5)
	wg.Add(2)
	go func(){
		defer wg.Done()
		time.Sleep(time.Second*5)
		intChan <- 10
	}()
	go func(){
		defer wg.Done()
		time.Sleep(time.Second*2)
		stringChan <- "1sd0"
	}()
	//select随机选择管道执行，stringChan阻塞时间短，则先执行
	select {
		case v:= <- intChan:
			fmt.Println("intChan:",v)
		case v:= <- stringChan:
			fmt.Println("stringChan:",v)
		default: //防止select被阻塞
			fmt.Println("Default")
	}
	wg.Wait()
}