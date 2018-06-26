package main

import (
	"net"
	"log"
	"fmt"
	"bytes"
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
				data, err := readScoketData(conn)
				if err != nil {
					break
				}
				fmt.Println(data)

				//4.返回给客户端数据
				conn.Write([]byte("echo"+data+"\t"))
			}
		}(conn)
	}


}

func readScoketData(conn net.Conn) (string, error) {
	//3读去服务端的响应信息read（）
	data := [1]byte{}
	buffer := bytes.Buffer{}
	for {
		_, err := conn.Read(data[:])
		if err != nil {
			fmt.Println(err)
			return "",err
		}
		if string(data[0]) == "\t" {
			break
		}
		buffer.Write(data[:])
	}

	return buffer.String(),nil
}
//比如客户端发送“你好"
//服务端返回”echo“+”你好“
