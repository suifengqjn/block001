package main

import "fmt"

func main() {
	/*
	pointer，指针
		概念：存储了另一个变量的内存地址的变量。

	指针的类型：*int,*float,*string,*array,*struct
	指针中存储的数据的类型:int,float,string,array,struct
	指针中存储的数据的地址：指针中存储的数值
	指针自己的地址
	 */
	//1.定义一个int类型的变量
	a := 10
	fmt.Println("a的数值是：", a)     //10
	fmt.Printf("a的类型是：%T\n", a)  //int
	fmt.Printf("a的地址是：%p\n", &a) //0xc42001c0c8

	//2.创建一个指针变量，用于存储变量a的地址
	var p1 *int
	fmt.Println(p1) //nil 空指针
	p1 = &a
	fmt.Println("p1的数值是：", p1) //p1中存储的数值，a的地址
	fmt.Println("p1的数值是变量的地址，该地址存储的数据：", *p1)
	fmt.Printf("p1自己的地址：%p\n", &p1)

	//3.操作变量，更改数值,并不会更改地址
	a = 100
	fmt.Println(a)
	fmt.Printf("%p\n", &a)

	//4.操作指针改变变量的数值
	*p1 = 200
	fmt.Println(a)

	//5.指针的指针
	var p2 **int
	fmt.Println(p2)
	p2 = &p1
	fmt.Println("p2的数值：",p2)//p1的地址
	fmt.Println("p2的数值，p1的地址，对应的数据：",*p2)//a的地址
	fmt.Println("p2中存储的地址对应的数值，在对应的数值：",**p2)

	fmt.Printf("%T,%T,%T\n",a,p1,p2)//int, *int, **int
}
/*
练习1：
创建变量，创建指针，操作指针更改变量，指针的指针。。
 */
