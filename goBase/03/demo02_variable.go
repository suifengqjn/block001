package main

import "fmt"

var b = 200//全局变量
var c int = 300
//d := 400 //语法错误

func main()  {
	/*
	局部变量：函数内部定义的变量，该变量的作用范围
	全局变量：函数外部，整个作用域
	 */
	 a := 100
	 fmt.Println(a)
	 fmt.Println(b)
}

func fun(){
	//fmt.Println(a)
	fmt.Println(b)
}
