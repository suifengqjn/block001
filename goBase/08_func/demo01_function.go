package main

import "fmt"

func main() {
	/*
	函数：
		1.概念：就是一段具有独立功能代码。可以被多次调用。

		2.意义：
			A：避免重复代码
			B：便于扩展程序的功能

		3.使用
			step1：函数的定义，也叫做声明
			step2：函数的调用，函数可以多次调用。调用函数位置叫调用处

		4.语法格式
			A：声明的语法格式
				func funcName(parametername type1, parametername type2) (output1 type1, output2 type2) {
					//这里是处理逻辑代码
					//返回多个值
					return value1, value2
				}
				func：关键字，专门用于定义函数
					var，const
				funcName：函数名，就是一个标识符，要满足标识符的命名规范
				()：是函数和方法的标志
					[],数组切片，{},

				parametername type1, parametername type2里的叫形参列表：形式参数。用于接收调用处传来的数据
				output1 type1, output2 type2：函数的返回值，调用函数后，返回给调用处的结果
			B：调用的语法格式
				通过函数名(实参列表)：实际参数

		5.注意事项
			A:函数必须先定义，再调用，
				如果不定义，undefined: getSum
				如果不调用，函数的定义没有意义了，白写了
			B：函数名不能冲突
			C：main()作为程序的入口，由系统自动调用，其他函数由程序中手动调用。
	 */

	//求1-10的和
	//调用函数
	getSum()

	fmt.Println("helloworld...")
	//求1-10的和
	getSum()

	fmt.Println("abc....")
	//....
	//求1-10的和
	getSum()
}

//1.定义一个函数，用于求1-10的和
func getSum(){
	//函数的代码：函数体
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1-100的和：", sum)
}
