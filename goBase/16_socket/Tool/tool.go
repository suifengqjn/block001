package Tool

import (
	"log"
	"fmt"
	"net"
	"bytes"
)

func CheckErr(err error)  {
	if err!= nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}
}

func ReadSocketData(conn net.Conn)(string, error)  {

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
