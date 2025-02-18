package main
import(
	"fmt"
	"net"
)
func process(conn net.Conn){
	defer conn.Close()
	for{
		//创建一个切片，将读取的切片放入切片
		buf := make([]byte,1024)
		n,err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:n]))
	}
}
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
		conn,err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("客户端的等待失败err2",err2)
		}else{
			fmt.Printf("连接成功，%v\n",conn.RemoteAddr().String())
		}
		//准备一个协程，处理客户端请求
		go process(conn)//不同客户请求连接不同
	}
}