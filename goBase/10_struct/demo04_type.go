package main

import "fmt"

type s1 struct {

}

//定义一个myint类型，
type my_int int

type my_fun func (int,int)(int,string)

func main() {
	/*
	var 变量名 数据类型
	const 常量名 数据类型=赋值
	func 函数名 函数体{}
	type 结构体名 struct{}

	type：用于定义数据类型
	 */
	 var a int
	 a = 100
	 var b my_int
	 b = 200
	 fmt.Println(a,b)
	 //a = b //cannot use b (type my_int) as type int in assignment

}
