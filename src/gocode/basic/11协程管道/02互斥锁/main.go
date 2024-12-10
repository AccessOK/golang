package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
//互斥锁
var lock sync.Mutex
var totalnum int
func add(){
	defer wg.Done()
	for i:= 0 ; i<100 ; i++ {
		lock.Lock()
		totalnum=totalnum+1
		lock.Unlock()
	}
}
func sub(){
	defer wg.Done()
	for i:= 0 ; i<100 ; i++ {
		lock.Lock()
		totalnum=totalnum-1
		lock.Unlock()
	}
}
func main(){
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println("end=",totalnum)
}