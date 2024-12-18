package main
import "fmt"
func main(){
	fmt.Println("--------------------------字符类型------------------------------------")
	var c1 byte = 'a'
	fmt.Println(c1)
	//golang的字符对应的使用时UTF-8编码(Unicode时对应的字符集，UTF-8是Unicode的其中的一种编码方案)
	var c2 byte = 'A'
	fmt.Printf("%c",c2)
	fmt.Println("--------------------------转义字符------------------------------------")
	// \b为退格
	fmt.Println("aaaa\bbbb")
	// \r将光标回到本行的开头，后续输入替换原有的字符
	fmt.Println("aaaa\rbbb")
	// \t制表符，每八位为一个制表位
	fmt.Println("aaaaaaa\tbbb")
	// \" \' \\原样输出
	fmt.Println("aaaaaaa\"bbb")
	fmt.Println("--------------------------布尔类型------------------------------------")
	//bool类型数据只运行true和false，之战1个字节
	var flag bool = true
	fmt.Println(flag)
	fmt.Println("--------------------------字符串类型------------------------------------")
	//定义字符串
	var s1 string = "hello golang"
	fmt.Println(s1)
	//字符串是不可变的
	var s2 string = "abc"
	fmt.Println(s2)
	//如果字符串中没有特殊字符，字符串的表示形式使用双引号
	//如果字符转中有特殊字符，字符串的表现形式使用反引号 ``
	var s3 string = `
	func main() { //程序入口
		fmt.Println("Hello, 世界")
	}
	`
	fmt.Println(s3)
	//字符串拼接
	var s4 string = "abc" + "def"
	s4 += "ghij"
	fmt.Println(s4)
	//字符串过长的时候，注意 + 保留在上一行
}