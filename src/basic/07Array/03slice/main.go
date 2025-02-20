package main
import "fmt"
func main(){
	//切片 对数组连续片段的引用
	var intarr [6]int = [6]int{2,4,5,1,4,5}
	//切片构建在数组之上,在数组上切出的片段
	var slice1 []int = intarr[2:5]
	fmt.Println(slice1)
	slice2 := intarr[1:4]
	fmt.Println(slice2)
	//切片元素个数
	fmt.Println(len(slice1))
	//切片容量
	fmt.Println(cap(slice1))
	//切片内容修改，源数组数据也修改
	slice1[0] = 99
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(intarr)
	//make 创建切片,1,切片类型，切片长度，切片容量
	slice3 :=make([]int,4,20)
	fmt.Println(slice3)
	slice3[0] = 98
	fmt.Println(slice3)
	//遍历切片
	for i := 0;i < len(slice3) ; i++ {
		fmt.Println(slice3[i])
	}
	for i,value := range slice3 {
		fmt.Printf("%d %d", i , value)
	}
	fmt.Println("-----------------------")
	//切片不可以直接使用，需要让其引用到一个数组或者使用make一个空间容切片使用
	//切片使用不能越界
	//切片可以继续切片
	//切片可以动态增长:创建一个新数组在追加数据，之前的数据不变
	intarr1 := append(slice3,88,50)
	fmt.Println(intarr1)
	slice3 = append(slice1,slice2...)
	fmt.Println(slice3)
	//底层的新数组，不能直接维护，只能通过切片进行间接维护
	var intarr2 []int = []int{1,3,5,7,4,2}
	var intarr3 []int = make([]int,10)
	//拷贝
	copy(intarr3,intarr2) // 将a中对应数组中元素内容复制到b中对应的数组
	fmt.Println(intarr3)
}