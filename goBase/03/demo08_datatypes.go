package main

import "fmt"

func main()  {
	/*
	数据类型：
		基本数据类型：
			数值类型：
				整数(int)，浮点(float)，复数(complex)
			布尔类型:bool
				数值：true，false
			字符串类型
		复合数据类型：
	 */
	 //1.定义布尔类型的变量
	 var b1 bool
	 b1 = true
	 b2 := false
	 //var b3 bool = 100
	 fmt.Println(b1,b2)
	 fmt.Printf("%T,%t\n",b1,b1)

	 //2.整数：
	 var i1 int8 = 100 //[-128,127]
	 fmt.Println(i1)

	 var i2 uint8 = 255 //[0,255]
	 fmt.Println(i2)

	 var i3 int = 1000
	 fmt.Println(i3)

	 var i4 byte = 200
	 fmt.Println(i4)

	 var i8 uint8
	 i8 = i4
	 fmt.Println(i8)

	 //var i5 int8 = 10
	 //var i6 int64 = 10
	 //var i7 int = 10
	 //i7 = i6

	 // 3.定义浮点类型的数据
	 var f1 float32
	 f1 = 3.14
	 var f2 float64
	 f2 = 3.14
	 fmt.Println(f1,f2)
	 fmt.Printf("%T,%.2f\n",f1,f1)

	 //4.定义字符串 ：多个byte的集合。
	 //理解为一个字符序列



}
