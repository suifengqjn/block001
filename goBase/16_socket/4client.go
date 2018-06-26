package main

import (
	"net"
	"log"
	"time"
	"fmt"
	"bytes"
)

func main()  {
	//1.连接服务端dial（）
	conn, err := net.Dial("tcp", "10.0.151.139:12345")
	if err != nil {
		log.Fatal()
	}
	defer conn.Close()


	for i:=0; i < 10; i++ {
		time.Sleep(time.Second)
		//2通过连接写入数据write（）
		conn.Write([]byte("hello server     \t"))

		data := readScoketData(conn)
		fmt.Println(data)
	}

}

func readScoketData(conn net.Conn) string {
	//3读去服务端的响应信息read（）
	data := [1]byte{}
	buffer := bytes.Buffer{}
	for {
		_, err := conn.Read(data[:])
		if err != nil {
			log.Fatal(err)
		}
		if string(data[0]) == "\t" {
			break
		}
		buffer.Write(data[:])
	}

	return buffer.String()
}
