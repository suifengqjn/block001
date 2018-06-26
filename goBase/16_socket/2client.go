package main

import (
	"net"
	"log"
	"fmt"
)

func main()  {
	//1.连接服务端dial（）
	conn, err := net.Dial("tcp", "10.0.151.139:12345")
	if err != nil {
		log.Fatal()
	}
	defer conn.Close()

	for i:=0; i < 10; i++ {
		//2通过连接写入数据write（）
		conn.Write([]byte("你好服务端"))

		//3读去服务端的响应信息read（）
		data := [512]byte{}
		n, err := conn.Read(data[:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data[:n]))
	}
}
