package main

import (
	"os"
	"fmt"
)

func main() {
	/*
	seek(offset,whence),设置指针光标的位置
	第一个参数：偏移量
	第二个参数：如何设置
		0：距离文件开头offset个字节
		1：距离光标当前的位置
		2：距离文件末尾

	随机读取文件：
		可以设置指针光标的位置


	bs : [1024]bs
	10次:10240

	临时文件：10240
	 */

	 file,_:=os.OpenFile("/home/ruby/文档/pro/aa.txt",os.O_RDWR,0)
	 bs :=[]byte{0}

	 file.Read(bs)
	 fmt.Println(string(bs))

	 file.Seek(4,0)
	 file.Read(bs)
	 fmt.Println(string(bs))
	 file.Seek(2,0)
	 file.Read(bs)
	 fmt.Println(string(bs))

	 file.Seek(3,1)
	 file.Read(bs)
	 fmt.Println(string(bs))

	 file.Seek(0,2)
	 file.WriteString("ABC")

}
