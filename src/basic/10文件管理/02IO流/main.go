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
	//写入文件
	file1,err1 := os.OpenFile("D:/AccessOK/code/golang/src/gocode/basic/10文件管理/test.txt",os.O_RDWR|os.O_APPEND,0666)
	if err1 != nil {
		fmt.Println("打开失败")
		return
	}
	defer file1.Close()
	writer := bufio.NewWriter(file1)
	for i := 0; i < 10; i++ {
		writer.WriteString("\n你好，这是输入流")
	}
	writer.Flush()//流数据在缓冲区中，刷新，将缓冲区中的内容保存到磁盘
	//文件复制
	//源文件
	content1,err1 := ioutil.ReadFile("D:/AccessOK/code/golang/src/gocode/basic/10文件管理/test.txt")
	if err != nil {
		fmt.Println("读取失败")
		return
	}
	//目标文件
	err2 := ioutil.WriteFile("D:/AccessOK/code/golang/src/gocode/basic/10文件管理/test1.txt",content1,0666)
	if err2 != nil{
		fmt.Println("写入失败！")
	}
}