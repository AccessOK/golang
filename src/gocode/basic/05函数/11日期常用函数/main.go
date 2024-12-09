
package main
import ("fmt"
		"time"
)
func main(){
	//获取日期
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("%v\n",now.Year())
	//日期格式化
	//这个参数字符串的各个数字必须时固定的，必须这样写
	fmt.Println(now.Format("2012/02/03"))
}