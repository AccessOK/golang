package main
import "fmt"
func main(){
	//map 键值对映射
	//只申明map内存是没有分配空间
	var a map[int]string
	//给map分配内存
	a = make(map[int]string,10)
	//赋值
	a[23212] = "hello"
	a[23212] = "hellosd"
	fmt.Println(a)
	//map的key和value是无序的
	//map的key不能重复，保存最后一次值，value可以重复
	b := make(map[int]string)
	b[23] = "qwe"
	b[343] = "weqw12"
	fmt.Println(b)
	c := map[int]string{
		123:"23123",
		23:"2312aeqw",
	}
	fmt.Println(c)
	//修改map的值
	c[23] = "qweqweqwe"
	fmt.Println(c[23])
	//删除map
	delete(c,23)
	fmt.Println(c)
	//查找
	value,flag := c[23]
	fmt.Printf("value=%s,flag=%t\n",value,flag)
	//map只支持for range便利
	for k,v := range b {
		fmt.Printf("%d,%s\n",k,v)
	}
	
}