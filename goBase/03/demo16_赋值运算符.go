package main

import "fmt"

func main()  {
	/*
	赋值运算符：=,+=,-=,*=,/=,%=,<<=,>>=,....
		a = 3,将=右侧的数值，赋值给=左侧的变量
	 */

	 var a int
	 a = 3
	 fmt.Println(a)//3

	 a += 2 //相当于a = a + 2
	 fmt.Println(a)//5

	 a /= 3 //a =a /3
	 fmt.Println(a)//1

	 a %= 1 //
	 fmt.Println(a)
}
