package main

import "fmt"

func main() {
	/*
	函数的参数：
		定义函数：形式参数，简称形参，接收调用处传入的数据的变量
			函数中，某些变量的数值，无法确定，需要由调用处传入数据。
		调用函数：实际参数，简称实参，调用函数的时候，给形参赋值的实际的数据

	函数调用：
		通过函数名，实参严格匹配形参。
			实参必须和形参一一对应：个数，类型，顺序
	 */
	 //1.求1-10的和
	 getSum(10)

	 //2.求1-20的和
	 getSum(20)
	 //3.求1-100的和
	 getSum(100)

	 fun1(20,10)

	 //fun2(3.14,"hello",100)
}

func getSum(n int){//需要变量n，n的值函数中无法确定
	sum := 0
	for i:=1;i<=n;i++{
		sum += i
	}
	fmt.Printf("1-%d的和：%d\n",n,sum)
}

func fun1(a int,b int){
	sum := 0
	sum = a + b
	fmt.Printf("%d + %d = %d \n",a, b, sum)
}

//以上代码可以简写
func fun3(a, b int){

}


func fun2(a int,b float64, c string){

}

func fun4(a,b,c int, s1,s2,s3 string){

}



