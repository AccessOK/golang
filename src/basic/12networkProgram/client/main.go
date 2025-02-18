package main
import(
	"fmt"
	"net"
	"os"
	"bufio"
)
func main(){
	fmt.Println("Client:")
	//调用Dial函数，发送请求
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err!= nil {
		fmt.Println("客户端连接失败:",err)
		return
	}else{
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		fmt.Println("客户端连接成功")
		//通过客户端发送单行数据
		reader := bufio.NewReader(os.Stdin)//os.Stdin终端标准输入
		//从终端读取一行用户输入的信息
		str,err1 := reader.ReadString('\n')
		if err1 != nil {
			fmt.Println("终端输入失败，err1=",err1)
		}
		n,err2 := conn.Write([]byte(str))
		if err2 != nil {
			fmt.Println("发送数据失败，err2=",err2)
		}else{
			fmt.Println("发生成功:",n)
		}
	}
}