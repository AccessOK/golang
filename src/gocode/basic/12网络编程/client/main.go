package main
import(
	"fmt"
	"net"
)
func main(){
	fmt.Println("Client:")
	//调用Dial函数，发送请求
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err!= nil {
		fmt.Println("客户端连接失败:",err)
		return
	}
	fmt.Println("客户端连接成功")
}