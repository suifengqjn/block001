package main

import "fmt"

//num2 := 100 //全局变量，不能省略var，不能使用简短的定义方式
var num2 = 1000//定义全局变量，大大


func main() {
	/*
	作用域：变量的可以使用的范围。
		局部变量：函数内生命的变量叫局部变量。作用域是函数所定义的部分
		全局变量：函数外声明的变量叫全局变量。作用域是整个文档。所有的函数都可以使用，而且是共享这一份全局变量。
	 */
	 n := 100
	 fmt.Println(n)//?

	 if a := 10;a >0{//a是if里定义的，那么作用域是if中
	 	fmt.Println(a)
	 	fmt.Println(n)
	 }
	 //fmt.Println(a)//无法使用a
	 fmt.Println(n)

	 if b := 20;n >0{
	 	n := 100
	 	fmt.Println(b)
	 	fmt.Println(n)
	 }
	 //fmt.Println(b)
	 fmt.Println(n)
	 {
	 	m := 200
	 	fmt.Println(m)
	 }
	 //fmt.Println(m)
	 fmt.Println("main:",num2)//1000

	 fun1()
	 fun2()
	 fmt.Println("main:",num2)//2000
}


func fun1(){
	fmt.Println("fun1...")
	num1 := 100//东北大哥
	fmt.Println(num1)
	num2 = 2000
	fmt.Println("fun1:",num2)//2000
}

func fun2(){
	fmt.Println("fun2...")
	//fmt.Println(num1)
	fmt.Println("fun2:",num2)//2000
}
