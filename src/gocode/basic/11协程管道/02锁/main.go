package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var totalnum int
func add(){
	defer wg.Done()
	for i:= 0 ; i<10000 ; i++ {
		totalnum=totalnum+1
	}
}
func sub(){
	defer wg.Done()
	for i:= 0 ; i<10000 ; i++ {
		totalnum=totalnum-i
	}
}
func main(){
	wg.Add(1)
	go add()
	go sub()
	wg.Wait()
	fmt.Println("end=",totalnum)
}