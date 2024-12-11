package main
import(
	"fmt"
	"net"
)
func main(){
	fmt.Println("Client:")
	//调用Dial函数，发送请求
	conn,err := net.Dial("tcp","baidu.com:80")
	if err!= nil {
		fmt.Println("客户端连接失败:",err)
		return
	}else{
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		fmt.Println("客户端连接成功")
	}
}