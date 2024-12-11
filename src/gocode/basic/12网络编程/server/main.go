package main
import(
	"fmt"
	"net"
)
func main(){
	fmt.Println("Server:")
	//调用Dial函数，进行监听
	listen,err1 := net.Listen("tcp","127.0.0.1:8888")
	if err1 != nil {
		fmt.Println("监听失败:",err1)
		return
	}
	for{
		//循环等待客户端连接，运行之后可以通过浏览器访问127.0.0.1:8888
		con,err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("客户端的等待失败err2",err2)
		}else{
			fmt.Printf("连接成功，%v",con.RemoteAddr().String())
		}
	}
}