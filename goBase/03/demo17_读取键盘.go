package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		fmt包：输入输出：I/O
			打印：输出
				Print()
				Printf()
				Println()
			输入：接收键盘的输入，标准输入
	*/

	//var a int // 0
	//var b float64 //0.0-->0
	//var c string // ""
	//fmt.Println("请输入一个整数，一个浮点，一个字符串：")
	//fmt.Printf("%p\n",&a)
	/*
	 Scanln(),读一行的数据，赋值给对应的变量，通过地址获取数据
	 数据和数据之间，空格隔开，按照顺序一次赋值



	 Scanf(), 第一个参数是读取的规则，
	 "%d,%d"
	*/
	//fmt.Scanln(&a,&b,&c)//阻塞式，

	//fmt.Scanf("%d,%f,%s",&a,&b,&c)
	//fmt.Println("a:",a)
	//fmt.Println("b:",b)
	//fmt.Println("c:",c)

	// var a int
	// var b string
	// fmt.Scanf("%d, %s", &a, &b)
	// fmt.Println(a, b)

	//var m string
	//fmt.Println("请输入一个字符串：")
	//fmt.Scanln(&m)
	//fmt.Println(m)

	//reader := bufio.NewReader(os.Stdin)
	// data, _, _ := reader.ReadLine()   //切片--->字符串
	// fmt.Printf("%T,%v\n", data, data) //[]uint8,[104 101 108 108 111 119 111 114 108 100]
	// fmt.Println(string(data))         //helloworld

	// s1, _ := reader.ReadString('?')
	// fmt.Println("---", s1)

	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadSlice('\n')
	fmt.Printf("%T,%v\n", data, data)
	fmt.Println(string(data))

}
