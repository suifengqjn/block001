package main

import "fmt"

func main()  {
	/*
	定义字符串 ：
		概念：多个byte的集合。
	 			理解为一个字符序列
		语法：
			使用双引号：""
			"","a","abc","hello"
			也可以使用``
					区别于单引号
		'A'-->int32
		"A"-->string

	转义字符：\
		A：有一些字符，转义后具有特殊的含义
			\n,\r,\t换行

		B：有一些字符，有特殊含义，转义后作为普通的字符
			\",\'
	 */
	 var s1 string
	 s1 = "王二狗"
	 fmt.Printf("%T,%s\n",s1,s1)

	 s2 := `hello world`
	 fmt.Println(s2)

	 var i1 = 'A'
	 fmt.Println(i1)
	 fmt.Printf("%c,%q\n",i1,i1)

	 fmt.Print("hello\nworld\tmemeda")

	 fmt.Println() //
	 fmt.Println("\"helloworld\"")
	 fmt.Println(`helo"王二狗"噶哈呢`)
}
/*
练习1：每种数据类型，分别定义2个变量，打印输出数据类型和数值。
练习2：打印输出：面朝"大海,春暖"花开
 */
