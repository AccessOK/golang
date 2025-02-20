
package main
import ("fmt"
		"strconv"
		"strings"
)
func main(){
	//统计字符串长度
	str := "asdasd"
	fmt.Println(len(str))
	//对字符串进行遍历
	for i:=0 ; i<len(str) ; i++{
		fmt.Printf("%c",str[i])
		fmt.Println(str[i])
	}
	for i,value := range str{
		fmt.Printf("%c",value)
		fmt.Println(i)
	}
	//字符串转换整数
	num,_ :=strconv.Atoi("666")
	fmt.Println(num)
	//整数转换为字符串
	str1 :=strconv.Itoa(88)
	fmt.Println(str1)
	//统计一个字符串中有几个指定的子串
	count := strings.Count("asdqwesdqweqweqzxda","qwe")
	fmt.Println(count)
	//不区分大小写的字符串比较
	flag := strings.EqualFold("hello","HELlo")
	fmt.Println(flag)
	//不区分大小写的字符串比较
	flag1 := "hello"=="HELlo"
	fmt.Println(flag1)
	//返回字串在字符串第一次出现的索引值
	index := strings.Index("asdqwesdqweqweqzxda","qwe")
	fmt.Println(index)
	//字符串替换
	replace := strings.Replace("asdqwesdqweqweqzxda","qwe","123",-1)
	fmt.Println(replace)
	//按照指定的某个字符，将字符串拆分为数组
	mark := strings.Split("gsudqw-weqw","-")
	for i := 0 ; i < len(mark) ; i++{
		fmt.Println(mark[i])
	}
	//将字符串的字母进行大小写转换
	fmt.Println(strings.ToLower("Hello"))
	fmt.Println(strings.ToUpper("Hello"))
	//将字符串左右两边的空格去掉
	fmt.Println(strings.TrimSpace("    go      "))
	//判断字符串开头
	fmt.Println(strings.HasPrefix("https://","http"))

}