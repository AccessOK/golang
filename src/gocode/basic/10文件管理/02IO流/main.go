package main
import(
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"bufio"
)
func main(){
	//一次性读取文件
	content,err := ioutil.ReadFile("D:/AccessOK/code/golang/src/gocode/basic/10文件管理/test.txt")
	if err != nil {
		fmt.Println("读取出错：",err)
	} else{
		fmt.Println(string(content))
	}
	//带缓冲的流
	//打开文件
	file,err := os.Open("D:/AccessOK/code/golang/src/gocode/basic/10文件管理/test.txt")
	if err != nil {
		fmt.Println("文件打开失败，err=",err)
	}
	//当函数退出时，让file关闭，防止内存泄露
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str,err :=reader.ReadString('\n')//读取到一个换行
		if err == io.EOF { //io.EOF 表述已经读取到文件的结尾
			break
		}
		fmt.Println(str)
	}
	fmt.Println("ending")
}