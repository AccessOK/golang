package main
import "fmt"
func main(){
	//二维数组
	var arr [2][3]int
	fmt.Println(arr)
	arr = [2][3]int{{2,4,5},{5,7,8}}
	fmt.Println(arr)
	//长度为2的一维数组
	//一维数组的值为一个长度为3的一维数组
	fmt.Printf("arr的地址%p\n",&arr)
	fmt.Printf("arr的地址%p\n",&arr[0])
	fmt.Printf("arr的地址%p\n",&arr[0][0])
	//二维数组便利
	for i := 0; i< len(arr); i++{
		for j := 0; j< len(arr[i]) ; j++{
			fmt.Printf("%d",arr[i][j])
		}
	}
	fmt.Println()
	for i,value1 := range arr{
		for j,value2 := range value1{
			fmt.Printf("%d %d %d\n",i,j,value2)
		}
	}
}