package main

import "fmt"

func main()  {
	/*
		零值，也叫默认值，某个变量没有给定具体的数值时，默认值
		int：0
		float：0
		string：""
	 */
	 var a int // 整数，默认值是0
	 var b float64//浮点，默认值0.0-->0
	 var s string //字符串：""
	 var s2 []int//切片sliece，引用类型 nil
	 fmt.Println(a,b,s,s2) //[]
	 fmt.Println(s2 == nil)



}
