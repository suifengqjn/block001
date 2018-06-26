package main

import "fmt"

func main() {
	/*
	匿名：没有名字
		匿名对象：
		匿名函数：函数没有名字

	通常只能使用一次。定义的时候直接使用
	 */
	 fun1()//调用函数
	 fun2:=fun1
	 fun2()//调用函数
	 //匿名函数：没有名字的函数
	 func(){
	 	fmt.Println("我是一个匿名函数。。")
	 }()//直接调用

	 func(a,b int){//带参数的匿名函数
	 	fmt.Println("a,",a,",b,",b)
	 }(1,2)

	 res1:=func(a,b,c int)int{//带返回值的匿名函数
	 	return a+b+c
	 }(1,2,3)
	 fmt.Println(res1)

	 res2:=func(){
	 	fmt.Println("我也是一个匿名函数。。。")
	 }
	 fmt.Println(res2)//0x4880b0

	 res3:=fun1
	 fmt.Println(res3)//0x487e60
	 res2()

}

func fun1(){
	fmt.Println("我是fun1函数。。。")

}



