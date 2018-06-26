package main

import "fmt"

func main()  {
	/*
	关系运算符：结果是bool类型
		>,<,>=,<=,==
			==,比较数值是否相等
			!=,比较数值是否不等

	 */
	 a := 3
	 b := 5
	 c := 3
	 fmt.Println(a > b,a > c)
	 fmt.Println(a <= b,a <= c)

	 fmt.Println(a == b)
	 fmt.Println(a == c)
	 fmt.Println(a != b)
	 fmt.Println(a != c)
}
