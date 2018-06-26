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

	for {
		//2.接受客户端的连接请求lister
		conn, err := lister.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//关闭连接
		defer conn.Close()

		go func(conn net.Conn) {
			for {
				//3.读客户端数据
				data := [512]byte{}
				n, err := conn.Read(data[:])
				if err != nil {
					//log.Fatal(err)
					fmt.Println(err)
					break
				}
				fmt.Println(string(data[:n]))

				//4.返回给客户端数据
				conn.Write([]byte("echo"+string(data[:n])))
			}
		}(conn)
	}


}

//比如客户端发送“你好"
//服务端返回”echo“+”你好“
