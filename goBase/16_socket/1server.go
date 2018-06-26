package main

import (
	"net"
	"log"
	"fmt"
)

func main() {
	//1.监听端口，并且绑定好了该端口
	lister, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	lister = LimitListener(lister, 10)
	//2.接受客户端的连接请求lister
	conn, err := lister.Accept()
	if err != nil {
		log.Fatal(err)
	}
	//5.关闭连接
	defer conn.Close()

	//3.读客户端数据
	data := [512]byte{}
	n, err := conn.Read(data[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data[:n]))

	//4.返回给客户端数据
	conn.Write([]byte("echo"+string(data[:n])))


}

//比如客户端发送“你好"
//服务端返回”echo“+”你好“

// 设置主机请求最大并发数
func LimitListener(l net.Listener, n int) net.Listener {

	return LimitListener(l, n)
}