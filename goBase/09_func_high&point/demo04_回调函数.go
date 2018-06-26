package main

import "fmt"

func main() {
	/*
	高阶函数：
		根据go语言函数的数据类型的特点，可以将函数作为另一个函数的参数。
		fun1()
		fun2()
	将fun1函数作为fun2函数的参数：
		fun2函数：高阶函数
			接收一个函数作为参数的函数，叫做高阶函数
		fun1函数：回调函数，
			作为另一个函数的参数的函数，叫回调函数


	回调函数：
	 */
	fmt.Printf("%T\n", add)  //func(int, int) int
	fmt.Printf("%T\n", oper) //func(int, int, func(int, int) int) int

	res1 := add(1, 2)
	fmt.Println(res1)

	res2 := oper(10, 20, add) //res2=res
	fmt.Println(res2)

	res3 := oper(7, 4, sub)
	fmt.Println(res3)

	fun1 := func(a, b int) int { //a=m=10,b=n=4
		fmt.Println("a:", a, ",b:", b)
		return a * b //40
	}

	res4 := oper(10, 4, fun1) //40
	fmt.Println(res4)

	res5 := oper(100, 8, func(a, b int) int {
		if b == 0 {
			fmt.Println("除数不能为零")
			return 0
		}
		return a / b
	})
	fmt.Println(res5)
	//fmt.Println(10/0)

}

//类型：func(int,int)int
func add(a, b int) int { //a=10,b=20
	return a + b //30
}
func sub(a, b int) int {
	return a - b
}

//类型：func(int,int, func(int,int)int)int
func oper(m, n int, fun func(int, int) int) int { //m=100,n=8,fun
	fmt.Println(m, n, fun)
	res := fun(m, n) //res=40
	return res
}
