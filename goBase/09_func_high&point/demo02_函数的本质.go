package main

import "fmt"

func main() {
	/*
	函数的本质：作为一种复合类型的数据，看作是一种特殊的变量。

	函数名()--->将函数进行调用，函数中的代码全部都执行一遍，然后将return后的结果返回
	函数名--->函数名，指向函数体的代码在内存中地址(理解为)


	函数可以作为一种特殊类型的变量：
		整型：算数运算：加减乘除，打印输出。。
		字符串：可以获取单个字符，截取子串，遍历。。strings包。。
		数组
		切片
		map
			容器：存数据，取数据，遍历，更改。。。。
		函数：调用

	 */

	 //1.定义一个整形变量
	 a := 10
	 //运算：加减乘除。。。打印
	 fmt.Println("a,",a*10)

	 //2.定义一个数组,切片，map
	 b :=[4]int{1,2,3,4}
	 b[0] =100
	 fmt.Println(b)

	 //3.函数作为一个变量
	 fmt.Printf("%T\n",fun1) //func (int,int)
	 fmt.Println(fun1) //0x489f10

	 //4.直接定义一个函数类型的变量
	 var c func (int,int)
	 //var d string
	 //d = "hello"
	 c = fun1//0x489f10
	 fun1(10,20)
	 c(100,200)


	 res1:=fun2(1,2)//将fun2函数进行调用，将执行结果赋值给res1,相当于：res1 = a+b
	 fmt.Println(res1)//3

	 res2 := fun2//将fun2的值(是一个函数，所以值是一个地址)，赋值给res2,
	 fmt.Println(res2)//0x48a0b0
	 res3 :=res2(10,20)
	 fmt.Println(res3)//30
	 //res1()//cannot call non-function res1 (type int)


}
func fun2(a,b int)int{
	return a + b
}

func fun1(a, b int){//类型：func (int,int)
	fmt.Println("a:",a,"b:",b)
}


