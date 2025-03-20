package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup //只需定义无需赋值
func sayHello() {
	for i := 0; i < 20; i++ {
		fmt.Println("Hello, go", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	wg.Done()
}
func main() { //主线程
	//协程在单线程中，主线程结束，协程也会结束
	wg.Add(1)     //协程开始时+1
	go sayHello() //开启协程
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, sort16422 ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	//多个协程,匿名函数直接获取函数外部i的变量，i的值不确定
	time.Sleep(time.Second * 15)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() //协程执行完成-1
			fmt.Println("Hello, many ", strconv.Itoa(i))

		}()
	}
	//多协程，且固定参数
	time.Sleep(time.Second * 15)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("Hello,muti ", strconv.Itoa(n))
			wg.Done()
		}(i)
		// time.Sleep(time.Second)
	}
	//主线程一直阻塞，wg=0时结束
	wg.Wait()
}
