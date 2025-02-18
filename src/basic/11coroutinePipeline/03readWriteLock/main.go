package main
import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
//读写锁
var lock sync.RWMutex

func read(){
	defer wg.Done()
	lock.RLock() //如果只是读数据，锁不产生影响，但是读写同时发生则会有影响
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}
func write(){
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second*10)
	fmt.Println("修改数据成功")
	lock.Unlock()
}
func main(){
	wg.Add(6)
	for i:=0 ; i < 5 ; i++{
		go read()
	}
	go write() // 写的过程中，锁生效
	wg.Wait()
}