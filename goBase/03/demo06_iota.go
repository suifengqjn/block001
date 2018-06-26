package main

import (
	"fmt"
)

func main()  {
	/*
	iota：计数器，go底层自动维护的一个计数器
		当使用const定义常量组的时候，初始化一个iota，初始值为0。。
			每当在常量组中定义一个常量，那么该计数器就累加1.
		直到下一个const出现，清零
	 */
	 const( //iota计数器：0
	 	a = iota // 0
	 	b = iota //1
	 	c = iota //2
	 )
	 fmt.Println("a:",a)
	 fmt.Println("b:",b)
	 fmt.Println("c:",c)

	 //常量组中的常量，没有初始值，默认和上一个
	 const(
	 	d = iota //0
	 	e
	 	f
	 )
	 fmt.Println(d,e,f)



}
