package main
import "fmt"
// 键盘录入
func main(){
	//scanln换行时才停止扫描
	//scanf 按照指定格式录入
	//键盘录入学生的年龄，姓名
	fmt.Println("请输入年龄！")
	var age int
	//在Scann函数张，对地址中的值进行改变的时候，实际外面的age被影响了
	//录入类型需要匹配，底层会自动判定类型
	fmt.Scanln(&age)
	fmt.Println(age)

	var name string
	fmt.Println("请输入姓名！")
	fmt.Scanf("%s",&name)
	fmt.Println(name)

}